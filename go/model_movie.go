/*
 * OpenAPI Movielist
 *
 * This is a sample server Movie list server.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Movie - A movie from the database
type Movie struct {

	// unique movie id
	Id int64 `json:"id"`

	Title string `json:"title"`

	ReleaseDate string `json:"release_date,omitempty"`

	// original release language of the movie
	OriginalLanguage string `json:"original_language,omitempty"`

	Adult bool `json:"adult,omitempty"`

	PosterPath string `json:"poster_path,omitempty"`

	Overview string `json:"overview,omitempty"`

	VoteAverage float32 `json:"vote_average,omitempty"`
}
