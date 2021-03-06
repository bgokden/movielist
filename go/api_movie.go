/*
 * OpenAPI Movielist
 *
 * This is a sample server Movie list server.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A MovieApiController binds http requests to an api service and writes the service results to the http response
type MovieApiController struct {
	service MovieApiServicer
}

// NewMovieApiController creates a default api controller
func NewMovieApiController(s MovieApiServicer) Router {
	return &MovieApiController{service: s}
}

// Routes returns all of the api route for the MovieApiController
func (c *MovieApiController) Routes() Routes {
	return Routes{
		{
			"GetMovieById",
			strings.ToUpper("Get"),
			"/v1/movie/{movieId}",
			c.GetMovieById,
		},
		{
			"ListMovies",
			strings.ToUpper("Get"),
			"/v1/movies",
			c.ListMovies,
		},
	}
}

// GetMovieById - Find movie by ID
func (c *MovieApiController) GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId, err := parseIntParameter(params["movieId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.GetMovieById(movieId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// ListMovies - List Movies
func (c *MovieApiController) ListMovies(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListMovies()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
