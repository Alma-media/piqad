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

type LocationApiService service

/* 
LocationApiService
Retrival of a single location
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Location unique ID
 * @param optional nil or *LocationGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return LocationFullResponse
*/

type LocationGetOpts struct { 
	ApiKey optional.String
}

func (a *LocationApiService) LocationGet(ctx context.Context, uid string, localVarOptionals *LocationGetOpts) (LocationFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LocationFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/location"

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
			var v LocationFullResponse
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
LocationApiService
Pagination over locations
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *LocationSearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return LocationBaseResponse
*/

type LocationSearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *LocationApiService) LocationSearchGet(ctx context.Context, localVarOptionals *LocationSearchGetOpts) (LocationBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LocationBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/location/search"

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
			var v LocationBaseResponse
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
LocationApiService
Searching locations
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *LocationSearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Name" (optional.String) -  Location name
     * @param "EarthlyLocation" (optional.Bool) -  Whether it should be an earthly location
     * @param "FictionalLocation" (optional.Bool) -  Whether it should be a fictional location
     * @param "ReligiousLocation" (optional.Bool) -  Whether it should be a religious location
     * @param "GeographicalLocation" (optional.Bool) -  Whether it should be a geographical location
     * @param "BodyOfWater" (optional.Bool) -  Whether it should be a body of water
     * @param "Country" (optional.Bool) -  Whether it should be a country
     * @param "SubnationalEntity" (optional.Bool) -  Whether it should be a subnational entity
     * @param "Settlement" (optional.Bool) -  Whether it should be a settlement
     * @param "UsSettlement" (optional.Bool) -  Whether it should be a US settlement
     * @param "BajoranSettlement" (optional.Bool) -  Whether it should be a Bajoran settlement
     * @param "Colony" (optional.Bool) -  Whether it should be a colony
     * @param "Landform" (optional.Bool) -  Whether it should be a landform
     * @param "Landmark" (optional.Bool) -  Whether it should be a landmark
     * @param "Road" (optional.Bool) -  Whether it should be a road
     * @param "Structure" (optional.Bool) -  Whether it should be a structure
     * @param "Shipyard" (optional.Bool) -  Whether it should be a shipyard
     * @param "BuildingInterior" (optional.Bool) -  Whether it should be a building interior
     * @param "Establishment" (optional.Bool) -  Whether it should be a establishment
     * @param "MedicalEstablishment" (optional.Bool) -  Whether it should be a medical establishment
     * @param "Ds9Establishment" (optional.Bool) -  Whether it should be a DS9 establishment
     * @param "School" (optional.Bool) -  Whether it should be a school
     * @param "Mirror" (optional.Bool) -  Whether this location should be from mirror universe
     * @param "AlternateReality" (optional.Bool) -  Whether this location should be from alternate reality

@return LocationBaseResponse
*/

type LocationSearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Name optional.String
	EarthlyLocation optional.Bool
	FictionalLocation optional.Bool
	ReligiousLocation optional.Bool
	GeographicalLocation optional.Bool
	BodyOfWater optional.Bool
	Country optional.Bool
	SubnationalEntity optional.Bool
	Settlement optional.Bool
	UsSettlement optional.Bool
	BajoranSettlement optional.Bool
	Colony optional.Bool
	Landform optional.Bool
	Landmark optional.Bool
	Road optional.Bool
	Structure optional.Bool
	Shipyard optional.Bool
	BuildingInterior optional.Bool
	Establishment optional.Bool
	MedicalEstablishment optional.Bool
	Ds9Establishment optional.Bool
	School optional.Bool
	Mirror optional.Bool
	AlternateReality optional.Bool
}

func (a *LocationApiService) LocationSearchPost(ctx context.Context, localVarOptionals *LocationSearchPostOpts) (LocationBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LocationBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/location/search"

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
	if localVarOptionals != nil && localVarOptionals.EarthlyLocation.IsSet() {
		localVarFormParams.Add("earthlyLocation", parameterToString(localVarOptionals.EarthlyLocation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FictionalLocation.IsSet() {
		localVarFormParams.Add("fictionalLocation", parameterToString(localVarOptionals.FictionalLocation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReligiousLocation.IsSet() {
		localVarFormParams.Add("religiousLocation", parameterToString(localVarOptionals.ReligiousLocation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GeographicalLocation.IsSet() {
		localVarFormParams.Add("geographicalLocation", parameterToString(localVarOptionals.GeographicalLocation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BodyOfWater.IsSet() {
		localVarFormParams.Add("bodyOfWater", parameterToString(localVarOptionals.BodyOfWater.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SubnationalEntity.IsSet() {
		localVarFormParams.Add("subnationalEntity", parameterToString(localVarOptionals.SubnationalEntity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Settlement.IsSet() {
		localVarFormParams.Add("settlement", parameterToString(localVarOptionals.Settlement.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UsSettlement.IsSet() {
		localVarFormParams.Add("usSettlement", parameterToString(localVarOptionals.UsSettlement.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BajoranSettlement.IsSet() {
		localVarFormParams.Add("bajoranSettlement", parameterToString(localVarOptionals.BajoranSettlement.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Colony.IsSet() {
		localVarFormParams.Add("colony", parameterToString(localVarOptionals.Colony.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Landform.IsSet() {
		localVarFormParams.Add("landform", parameterToString(localVarOptionals.Landform.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Landmark.IsSet() {
		localVarFormParams.Add("landmark", parameterToString(localVarOptionals.Landmark.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Road.IsSet() {
		localVarFormParams.Add("road", parameterToString(localVarOptionals.Road.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Structure.IsSet() {
		localVarFormParams.Add("structure", parameterToString(localVarOptionals.Structure.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Shipyard.IsSet() {
		localVarFormParams.Add("shipyard", parameterToString(localVarOptionals.Shipyard.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BuildingInterior.IsSet() {
		localVarFormParams.Add("buildingInterior", parameterToString(localVarOptionals.BuildingInterior.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Establishment.IsSet() {
		localVarFormParams.Add("establishment", parameterToString(localVarOptionals.Establishment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MedicalEstablishment.IsSet() {
		localVarFormParams.Add("medicalEstablishment", parameterToString(localVarOptionals.MedicalEstablishment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ds9Establishment.IsSet() {
		localVarFormParams.Add("ds9Establishment", parameterToString(localVarOptionals.Ds9Establishment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.School.IsSet() {
		localVarFormParams.Add("school", parameterToString(localVarOptionals.School.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mirror.IsSet() {
		localVarFormParams.Add("mirror", parameterToString(localVarOptionals.Mirror.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AlternateReality.IsSet() {
		localVarFormParams.Add("alternateReality", parameterToString(localVarOptionals.AlternateReality.Value(), ""))
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
			var v LocationBaseResponse
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
