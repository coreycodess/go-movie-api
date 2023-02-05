package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}

var movies = []movie{
	{ID: "1", Title: "Scarface", Rating: 8.5},
	{ID: "2", Title: "Godfather", Rating: 9},
	{ID: "3", Title: "Harry Potter", Rating: 8.8},
}

func getMovies(c *gin.Context){
	c.IndentedJSON(http.StatusOK, movies)
}

func createMovie(c *gin.Context){
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		return 
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func deleteMovie(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	movies = append(movies[:id], movies[id+1:]...)

	c.IndentedJSON(http.StatusAccepted, "deleted")
}

func main() {
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.POST("/movies/create", createMovie)
	router.DELETE("/movies/delete/:id", deleteMovie)
	router.Run("localhost:8080")
}