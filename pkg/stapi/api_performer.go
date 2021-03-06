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

type PerformerApiService service

/* 
PerformerApiService
Retrival of a single performer
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Performer unique ID
 * @param optional nil or *PerformerGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return PerformerFullResponse
*/

type PerformerGetOpts struct { 
	ApiKey optional.String
}

func (a *PerformerApiService) PerformerGet(ctx context.Context, uid string, localVarOptionals *PerformerGetOpts) (PerformerFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PerformerFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/performer"

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
			var v PerformerFullResponse
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
PerformerApiService
Pagination over performers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PerformerSearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return PerformerBaseResponse
*/

type PerformerSearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *PerformerApiService) PerformerSearchGet(ctx context.Context, localVarOptionals *PerformerSearchGetOpts) (PerformerBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PerformerBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/performer/search"

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
			var v PerformerBaseResponse
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
PerformerApiService
Searching performers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PerformerSearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Name" (optional.String) -  Performer name
     * @param "BirthName" (optional.String) -  Performer birth name
     * @param "Gender" (optional.String) -  Performer gender
     * @param "DateOfBirthFrom" (optional.String) -  Minimal date the performer was born
     * @param "DateOfBirthTo" (optional.String) -  Maximal date the performer was born
     * @param "PlaceOfBirth" (optional.String) -  Place the performer was born
     * @param "DateOfDeathFrom" (optional.String) -  Minimal date the performer died
     * @param "DateOfDeathTo" (optional.String) -  Maximal date the performer died
     * @param "PlaceOfDeath" (optional.String) -  Place the performer died
     * @param "AnimalPerformer" (optional.Bool) -  Whether it should be an animal performer
     * @param "DisPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: Discovery
     * @param "Ds9Performer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: Deep Space Nine
     * @param "EntPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: Enterprise
     * @param "FilmPerformer" (optional.Bool) -  Whether it should be a performer that appeared in a Star Trek movie
     * @param "StandInPerformer" (optional.Bool) -  Whether it should be a stand-in performer
     * @param "StuntPerformer" (optional.Bool) -  Whether it should be a stunt performer
     * @param "TasPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: The Animated Series
     * @param "TngPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: The Next Generation
     * @param "TosPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: The Original Series
     * @param "VideoGamePerformer" (optional.Bool) -  Whether it should be a video game performer
     * @param "VoicePerformer" (optional.Bool) -  Whether it should be a voice performer
     * @param "VoyPerformer" (optional.Bool) -  Whether it should be a performer that appeared in Star Trek: Voyager

@return PerformerBaseResponse
*/

type PerformerSearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Name optional.String
	BirthName optional.String
	Gender optional.String
	DateOfBirthFrom optional.String
	DateOfBirthTo optional.String
	PlaceOfBirth optional.String
	DateOfDeathFrom optional.String
	DateOfDeathTo optional.String
	PlaceOfDeath optional.String
	AnimalPerformer optional.Bool
	DisPerformer optional.Bool
	Ds9Performer optional.Bool
	EntPerformer optional.Bool
	FilmPerformer optional.Bool
	StandInPerformer optional.Bool
	StuntPerformer optional.Bool
	TasPerformer optional.Bool
	TngPerformer optional.Bool
	TosPerformer optional.Bool
	VideoGamePerformer optional.Bool
	VoicePerformer optional.Bool
	VoyPerformer optional.Bool
}

func (a *PerformerApiService) PerformerSearchPost(ctx context.Context, localVarOptionals *PerformerSearchPostOpts) (PerformerBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PerformerBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/performer/search"

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
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarFormParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BirthName.IsSet() {
		localVarFormParams.Add("birthName", parameterToString(localVarOptionals.BirthName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Gender.IsSet() {
		localVarFormParams.Add("gender", parameterToString(localVarOptionals.Gender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateOfBirthFrom.IsSet() {
		localVarFormParams.Add("dateOfBirthFrom", parameterToString(localVarOptionals.DateOfBirthFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateOfBirthTo.IsSet() {
		localVarFormParams.Add("dateOfBirthTo", parameterToString(localVarOptionals.DateOfBirthTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlaceOfBirth.IsSet() {
		localVarFormParams.Add("placeOfBirth", parameterToString(localVarOptionals.PlaceOfBirth.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateOfDeathFrom.IsSet() {
		localVarFormParams.Add("dateOfDeathFrom", parameterToString(localVarOptionals.DateOfDeathFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateOfDeathTo.IsSet() {
		localVarFormParams.Add("dateOfDeathTo", parameterToString(localVarOptionals.DateOfDeathTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlaceOfDeath.IsSet() {
		localVarFormParams.Add("placeOfDeath", parameterToString(localVarOptionals.PlaceOfDeath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AnimalPerformer.IsSet() {
		localVarFormParams.Add("animalPerformer", parameterToString(localVarOptionals.AnimalPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DisPerformer.IsSet() {
		localVarFormParams.Add("disPerformer", parameterToString(localVarOptionals.DisPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ds9Performer.IsSet() {
		localVarFormParams.Add("ds9Performer", parameterToString(localVarOptionals.Ds9Performer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EntPerformer.IsSet() {
		localVarFormParams.Add("entPerformer", parameterToString(localVarOptionals.EntPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilmPerformer.IsSet() {
		localVarFormParams.Add("filmPerformer", parameterToString(localVarOptionals.FilmPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StandInPerformer.IsSet() {
		localVarFormParams.Add("standInPerformer", parameterToString(localVarOptionals.StandInPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StuntPerformer.IsSet() {
		localVarFormParams.Add("stuntPerformer", parameterToString(localVarOptionals.StuntPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TasPerformer.IsSet() {
		localVarFormParams.Add("tasPerformer", parameterToString(localVarOptionals.TasPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TngPerformer.IsSet() {
		localVarFormParams.Add("tngPerformer", parameterToString(localVarOptionals.TngPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TosPerformer.IsSet() {
		localVarFormParams.Add("tosPerformer", parameterToString(localVarOptionals.TosPerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VideoGamePerformer.IsSet() {
		localVarFormParams.Add("videoGamePerformer", parameterToString(localVarOptionals.VideoGamePerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VoicePerformer.IsSet() {
		localVarFormParams.Add("voicePerformer", parameterToString(localVarOptionals.VoicePerformer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VoyPerformer.IsSet() {
		localVarFormParams.Add("voyPerformer", parameterToString(localVarOptionals.VoyPerformer.Value(), ""))
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
			var v PerformerBaseResponse
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
