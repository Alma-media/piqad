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

type BookApiService service

/* 
BookApiService
Retrival of a single book
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Book unique ID
 * @param optional nil or *BookGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return BookFullResponse
*/

type BookGetOpts struct { 
	ApiKey optional.String
}

func (a *BookApiService) BookGet(ctx context.Context, uid string, localVarOptionals *BookGetOpts) (BookFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BookFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/book"

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
			var v BookFullResponse
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
BookApiService
Pagination over books
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *BookSearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return BookBaseResponse
*/

type BookSearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *BookApiService) BookSearchGet(ctx context.Context, localVarOptionals *BookSearchGetOpts) (BookBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BookBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/book/search"

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
			var v BookBaseResponse
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
BookApiService
Searching books
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *BookSearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Title" (optional.String) -  Book title
     * @param "PublishedYearFrom" (optional.Int32) -  Starting year the book was published
     * @param "PublishedYearTo" (optional.Int32) -  Ending year the book was published
     * @param "NumberOfPagesFrom" (optional.Int32) -  Minimal number of pages
     * @param "NumberOfPagesTo" (optional.Int32) -  Maximal number of pages
     * @param "StardateFrom" (optional.Float32) -  Starting stardate of book story
     * @param "StardateTo" (optional.Float32) -  Ending stardate of book story
     * @param "YearFrom" (optional.Int32) -  Starting year of book story
     * @param "YearTo" (optional.Int32) -  Ending year of book story
     * @param "Novel" (optional.Bool) -  Whether it should be a novel
     * @param "ReferenceBook" (optional.Bool) -  Whether it should be a reference book
     * @param "BiographyBook" (optional.Bool) -  Whether it should be a biography book
     * @param "RolePlayingBook" (optional.Bool) -  Whether it should be a role playing book
     * @param "EBook" (optional.Bool) -  Whether it should be an e-book
     * @param "Anthology" (optional.Bool) -  Whether it should be an anthology
     * @param "Novelization" (optional.Bool) -  Whether it should be novelization
     * @param "Audiobook" (optional.Bool) -  Whether it should be an audiobook
     * @param "AudiobookAbridged" (optional.Bool) -  Whether it should be an audiobook, abridged
     * @param "AudiobookPublishedYearFrom" (optional.Int32) -  Starting year the audiobook was published
     * @param "AudiobookPublishedYearTo" (optional.Int32) -  Ending year the audiobook was published
     * @param "AudiobookRunTimeFrom" (optional.Int32) -  Minimal audiobook run time, in minutes
     * @param "AudiobookRunTimeTo" (optional.Int32) -  Maximal audiobook run time, in minutes

@return BookBaseResponse
*/

type BookSearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Title optional.String
	PublishedYearFrom optional.Int32
	PublishedYearTo optional.Int32
	NumberOfPagesFrom optional.Int32
	NumberOfPagesTo optional.Int32
	StardateFrom optional.Float32
	StardateTo optional.Float32
	YearFrom optional.Int32
	YearTo optional.Int32
	Novel optional.Bool
	ReferenceBook optional.Bool
	BiographyBook optional.Bool
	RolePlayingBook optional.Bool
	EBook optional.Bool
	Anthology optional.Bool
	Novelization optional.Bool
	Audiobook optional.Bool
	AudiobookAbridged optional.Bool
	AudiobookPublishedYearFrom optional.Int32
	AudiobookPublishedYearTo optional.Int32
	AudiobookRunTimeFrom optional.Int32
	AudiobookRunTimeTo optional.Int32
}

func (a *BookApiService) BookSearchPost(ctx context.Context, localVarOptionals *BookSearchPostOpts) (BookBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue BookBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/book/search"

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
	if localVarOptionals != nil && localVarOptionals.NumberOfPagesFrom.IsSet() {
		localVarFormParams.Add("numberOfPagesFrom", parameterToString(localVarOptionals.NumberOfPagesFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NumberOfPagesTo.IsSet() {
		localVarFormParams.Add("numberOfPagesTo", parameterToString(localVarOptionals.NumberOfPagesTo.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.Novel.IsSet() {
		localVarFormParams.Add("novel", parameterToString(localVarOptionals.Novel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReferenceBook.IsSet() {
		localVarFormParams.Add("referenceBook", parameterToString(localVarOptionals.ReferenceBook.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BiographyBook.IsSet() {
		localVarFormParams.Add("biographyBook", parameterToString(localVarOptionals.BiographyBook.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RolePlayingBook.IsSet() {
		localVarFormParams.Add("rolePlayingBook", parameterToString(localVarOptionals.RolePlayingBook.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EBook.IsSet() {
		localVarFormParams.Add("eBook", parameterToString(localVarOptionals.EBook.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Anthology.IsSet() {
		localVarFormParams.Add("anthology", parameterToString(localVarOptionals.Anthology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Novelization.IsSet() {
		localVarFormParams.Add("novelization", parameterToString(localVarOptionals.Novelization.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Audiobook.IsSet() {
		localVarFormParams.Add("audiobook", parameterToString(localVarOptionals.Audiobook.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AudiobookAbridged.IsSet() {
		localVarFormParams.Add("audiobookAbridged", parameterToString(localVarOptionals.AudiobookAbridged.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AudiobookPublishedYearFrom.IsSet() {
		localVarFormParams.Add("audiobookPublishedYearFrom", parameterToString(localVarOptionals.AudiobookPublishedYearFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AudiobookPublishedYearTo.IsSet() {
		localVarFormParams.Add("audiobookPublishedYearTo", parameterToString(localVarOptionals.AudiobookPublishedYearTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AudiobookRunTimeFrom.IsSet() {
		localVarFormParams.Add("audiobookRunTimeFrom", parameterToString(localVarOptionals.AudiobookRunTimeFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AudiobookRunTimeTo.IsSet() {
		localVarFormParams.Add("audiobookRunTimeTo", parameterToString(localVarOptionals.AudiobookRunTimeTo.Value(), ""))
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
			var v BookBaseResponse
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
