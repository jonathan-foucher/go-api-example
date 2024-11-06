package handlers

import (
 	"fmt"
 	"encoding/json"
 	"net/http"
 	"strconv"
	"go-api-example/models"
	"go-api-example/utils"
)

func HandleMovieRequest(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Println("Get all movies")
		res.Write([]byte ("Get all movies"))

	case "POST":
		movie := models.Movie{}
		json.NewDecoder(req.Body).Decode(&movie)
		fmt.Println("Post movie with id", movie.Id)
		fmt.Println("Post movie with title", movie.Title)
		fmt.Println("Post movie with release date", movie.ReleaseDate)
		res.Write([]byte (fmt.Sprintf("Post movie with id %d", movie.Id)))

	case "DELETE":
		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)
		movieId, err := strconv.Atoi(head)
		if err != nil {
			http.Error(res, fmt.Sprintf("Invalid movie id", head), http.StatusBadRequest)
			return
		}
		fmt.Println("Delete movie with id", movieId)
		res.Write([]byte (fmt.Sprintf("Delete movie with id %d", movieId)))

	default:
		http.Error(res, "Only GET, POST and DELETE are allowed", http.StatusMethodNotAllowed)
	}
}
