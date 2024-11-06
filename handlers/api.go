package handlers

import (
 	"net/http"
	"go-api-example/utils"
)

type ApiHandler struct {
}

func (h *ApiHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = utils.ShiftPath(req.URL.Path)
	if head == "api" {
		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)

		if head == "movies" {
			HandleMovieRequest(res, req)
		}
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}
