# go-movie-api

I created a very simple api that gets all the movies, creates movies and deletes movies based on the index in the array.

#### GET /movies
curl localhost:8080/movies

#### POST /movies/create
curl localhost:8080/movies/create -d @data.json

#### DELETE /movies/delete/:id
curl -X DELETE localhost:8080/movies/delete/1

---
## Software I Used:
* GO
---
