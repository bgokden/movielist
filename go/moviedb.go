package openapi

import (
	"os"
	"time"

	"github.com/goburrow/cache"
	"github.com/ryanbradynd05/go-tmdb"
)

// use this api key or use your own API key
// Check https://github.com/av150297/GOAPI/blob/master/main.go
const FALLBACK_APIKEY = "381193835c01836af08b2b8b05341ae5"

var tmdbAPI *tmdb.TMDb
var popularCache cache.LoadingCache
var movieCache cache.LoadingCache

func init() {

	config := tmdb.Config{
		APIKey:   getEnv("APIKEY", FALLBACK_APIKEY),
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)

	loadPopular := func(_ cache.Key) (cache.Value, error) {
		return GetMoviesFromTMDB()
	}
	popularCache = cache.NewLoadingCache(loadPopular,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(10*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)

	loadMovie := func(k cache.Key) (cache.Value, error) {
		return GetMovieFromTMDB(k.(int64))
	}

	movieCache = cache.NewLoadingCache(loadMovie,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(10*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetMoviesFromTMDB retrieves popular movie id and titles from TMDB
func GetMoviesFromTMDB() ([]Movie, error) {
	options := map[string]string{}
	moviePagedResults, err := tmdbAPI.GetMoviePopular(options)
	if err != nil {
		return nil, nil
	}
	movies := make([]Movie, 0, len(moviePagedResults.Results))
	for _, movieShort := range moviePagedResults.Results {
		movie := Movie{
			Id:    int64(movieShort.ID),
			Title: movieShort.Title,
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// GetMovieFromTMDB retrieves movie info from TMDB and return a movie
func GetMovieFromTMDB(movieID int64) (Movie, error) {
	options := map[string]string{}
	movieInfo, err := tmdbAPI.GetMovieInfo(int(movieID), options)
	if err != nil {
		return Movie{}, err
	}
	movie := Movie{
		Id:               int64(movieInfo.ID),
		Title:            movieInfo.Title,
		ReleaseDate:      movieInfo.ReleaseDate,
		OriginalLanguage: movieInfo.OriginalLanguage,
		Adult:            movieInfo.Adult,
		PosterPath:       movieInfo.PosterPath,
		Overview:         movieInfo.Overview,
		VoteAverage:      movieInfo.VoteAverage,
	}
	return movie, nil
}

// GetMovies returns the list of popular movies using cache
func GetMovies() ([]Movie, error) {
	value, err := popularCache.Get("popular")
	if err != nil {
		return nil, err
	}
	return value.([]Movie), nil
}

// GetMovie returns the movie info using cache
func GetMovie(movieID int64) (Movie, error) {
	value, err := movieCache.Get(movieID)
	if err != nil {
		return Movie{}, err
	}
	return value.(Movie), nil
}
