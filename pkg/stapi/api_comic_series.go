/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ComicSeriesApiService service

/* 
ComicSeriesApiService
Retrival of a single comic series
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Comic series unique ID
 * @param optional nil or *ComicSeriesGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return ComicSeriesFullResponse
*/

type ComicSeriesGetOpts struct { 
	ApiKey optional.String
}

func (a *ComicSeriesApiService) ComicSeriesGet(ctx context.Context, uid string, localVarOptionals *ComicSeriesGetOpts) (ComicSeriesFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ComicSeriesFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/comicSeries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("uid", parameterToString(uid, ""))
	if localVarOptionals != nil && localVarOptionals.ApiKey.IsSet() {
		localVarQueryParams.Add("apiKey", parameterToString(localVarOptionals.ApiKey.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ComicSeriesFullResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ComicSeriesApiService
Pagination over comic series
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ComicSeriesSearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return ComicSeriesBaseResponse
*/

type ComicSeriesSearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *ComicSeriesApiService) ComicSeriesSearchGet(ctx context.Context, localVarOptionals *ComicSeriesSearchGetOpts) (ComicSeriesBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ComicSeriesBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/comicSeries/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PageNumber.IsSet() {
		localVarQueryParams.Add("pageNumber", parameterToString(localVarOptionals.PageNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ApiKey.IsSet() {
		localVarQueryParams.Add("apiKey", parameterToString(localVarOptionals.ApiKey.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ComicSeriesBaseResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ComicSeriesApiService
Searching comic series
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ComicSeriesSearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Title" (optional.String) -  Comic series title
     * @param "PublishedYearFrom" (optional.Int32) -  Starting year the comic series was published
     * @param "PublishedYearTo" (optional.Int32) -  Ending year the comic series was published
     * @param "NumberOfIssuesFrom" (optional.Int32) -  Minimal number of issues
     * @param "NumberOfIssuesTo" (optional.Int32) -  Maximal number of issues
     * @param "StardateFrom" (optional.Float32) -  Starting stardate of comic series stories
     * @param "StardateTo" (optional.Float32) -  Starting stardate of comic series stories
     * @param "YearFrom" (optional.Int32) -  Starting year of comic series stories
     * @param "YearTo" (optional.Int32) -  Ending year of comic series stories
     * @param "Miniseries" (optional.Bool) -  Whether it should be a miniseries
     * @param "PhotonovelSeries" (optional.Bool) -  Whether it should be photonovel series

@return ComicSeriesBaseResponse
*/

type ComicSeriesSearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Title optional.String
	PublishedYearFrom optional.Int32
	PublishedYearTo optional.Int32
	NumberOfIssuesFrom optional.Int32
	NumberOfIssuesTo optional.Int32
	StardateFrom optional.Float32
	StardateTo optional.Float32
	YearFrom optional.Int32
	YearTo optional.Int32
	Miniseries optional.Bool
	PhotonovelSeries optional.Bool
}

func (a *ComicSeriesApiService) ComicSeriesSearchPost(ctx context.Context, localVarOptionals *ComicSeriesSearchPostOpts) (ComicSeriesBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ComicSeriesBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/comicSeries/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PageNumber.IsSet() {
		localVarQueryParams.Add("pageNumber", parameterToString(localVarOptionals.PageNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ApiKey.IsSet() {
		localVarQueryParams.Add("apiKey", parameterToString(localVarOptionals.ApiKey.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
		localVarFormParams.Add("title", parameterToString(localVarOptionals.Title.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PublishedYearFrom.IsSet() {
		localVarFormParams.Add("publishedYearFrom", parameterToString(localVarOptionals.PublishedYearFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PublishedYearTo.IsSet() {
		localVarFormParams.Add("publishedYearTo", parameterToString(localVarOptionals.PublishedYearTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NumberOfIssuesFrom.IsSet() {
		localVarFormParams.Add("numberOfIssuesFrom", parameterToString(localVarOptionals.NumberOfIssuesFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NumberOfIssuesTo.IsSet() {
		localVarFormParams.Add("numberOfIssuesTo", parameterToString(localVarOptionals.NumberOfIssuesTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StardateFrom.IsSet() {
		localVarFormParams.Add("stardateFrom", parameterToString(localVarOptionals.StardateFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StardateTo.IsSet() {
		localVarFormParams.Add("stardateTo", parameterToString(localVarOptionals.StardateTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.YearFrom.IsSet() {
		localVarFormParams.Add("yearFrom", parameterToString(localVarOptionals.YearFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.YearTo.IsSet() {
		localVarFormParams.Add("yearTo", parameterToString(localVarOptionals.YearTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Miniseries.IsSet() {
		localVarFormParams.Add("miniseries", parameterToString(localVarOptionals.Miniseries.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PhotonovelSeries.IsSet() {
		localVarFormParams.Add("photonovelSeries", parameterToString(localVarOptionals.PhotonovelSeries.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ComicSeriesBaseResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
