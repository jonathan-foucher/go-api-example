package handlers

import (
	"context"
 	"fmt"
 	"encoding/json"
 	"net/http"
 	"strconv"
  	db "go-api-example/database"
	"go-api-example/models"
	"go-api-example/utils"
)

func HandleMovieRequest(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		results, err := getAllMovies()
		res.Header().Set("Content-Type", "application/json")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
			return
		}
		json.NewEncoder(res).Encode(results)

	case "POST":
		movie := models.Movie{}
		json.NewDecoder(req.Body).Decode(&movie)
		err := SaveMovie(movie)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
		}

	case "DELETE":
		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)
		movieId, err := strconv.Atoi(head)
		if err != nil {
			http.Error(res, "Invalid movie id" + head, http.StatusBadRequest)
			return
		}
		DeleteMovie(int32(movieId))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
		}

	default:
		http.Error(res, "Only GET, POST and DELETE are allowed", http.StatusMethodNotAllowed)
	}
}

func getAllMovies() ([]models.Movie, error) {
	queries := db.New(db.GetDbConnection())

	fmt.Println("Get all movies")
	results, err := queries.GetMovies(context.Background())
	if err != nil {
		return nil, err
	}

	movies := make([]models.Movie, len(results))
    for i, result := range results {
        movies[i] = convertDbMovieToModelsMovie(result)
    }
	return movies, nil
}

func SaveMovie(movie models.Movie) error {
	queries := db.New(db.GetDbConnection())

	fmt.Println("Post movie with id", movie.Id)
	fmt.Println("Post movie with title", movie.Title)
	fmt.Println("Post movie with release date", movie.ReleaseDate)
	return queries.SaveMovie(context.Background(), db.SaveMovieParams{
		ID: movie.Id,
		Title: movie.Title,
		ReleaseDate: movie.ReleaseDate,
	})
}

func DeleteMovie(movieId int32) error {
	queries := db.New(db.GetDbConnection())

	fmt.Println("Delete movie with id", movieId)
	return queries.DeleteMovie(context.Background(), movieId)
}

func convertDbMovieToModelsMovie(dbMovie db.Movie) (models.Movie) {
	movie := models.Movie{
		Id: dbMovie.ID,
		Title: dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate,
	}
	return movie
}
