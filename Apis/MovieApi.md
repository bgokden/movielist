# MovieApi

All URIs are relative to *http://movielist.swagger.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMovieById**](MovieApi.md#getMovieById) | **GET** /movie/{movieId} | Find movie by ID
[**listMovies**](MovieApi.md#listMovies) | **GET** /movies | List Movies


<a name="getMovieById"></a>
# **getMovieById**
> Movie getMovieById(movieId)

Find movie by ID

    Returns a single movie

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **Long**| The id of the movie to retrieve | [default to null]

### Return type

[**Movie**](..//Models/Movie.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listMovies"></a>
# **listMovies**
> List listMovies()

List Movies

    Lists all movies

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](..//Models/Movie.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

