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

type TechnologyApiService service

/* 
TechnologyApiService
Retrival of a single technology
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Technology unique ID
 * @param optional nil or *TechnologyGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return TechnologyFullResponse
*/

type TechnologyGetOpts struct { 
	ApiKey optional.String
}

func (a *TechnologyApiService) TechnologyGet(ctx context.Context, uid string, localVarOptionals *TechnologyGetOpts) (TechnologyFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TechnologyFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/technology"

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
			var v TechnologyFullResponse
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
TechnologyApiService
Pagination over technology
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TechnologySearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return TechnologyBaseResponse
*/

type TechnologySearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *TechnologyApiService) TechnologySearchGet(ctx context.Context, localVarOptionals *TechnologySearchGetOpts) (TechnologyBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TechnologyBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/technology/search"

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
			var v TechnologyBaseResponse
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
TechnologyApiService
Searching technology
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TechnologySearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Name" (optional.String) -  Technology name
     * @param "BorgTechnology" (optional.Bool) -  Whether it should be a Borg technology
     * @param "BorgComponent" (optional.Bool) -  Whether it should be a Borg component
     * @param "CommunicationsTechnology" (optional.Bool) -  Whether it should be a communications technology
     * @param "ComputerTechnology" (optional.Bool) -  Whether it should be a computer technology
     * @param "ComputerProgramming" (optional.Bool) -  Whether it should be a technology related to computer programming
     * @param "Subroutine" (optional.Bool) -  Whether it should be a subroutine
     * @param "Database" (optional.Bool) -  Whether it should be a database
     * @param "EnergyTechnology" (optional.Bool) -  Whether it should be a energy technology
     * @param "FictionalTechnology" (optional.Bool) -  Whether it should be a fictional technology
     * @param "HolographicTechnology" (optional.Bool) -  Whether it should be a holographic technology
     * @param "IdentificationTechnology" (optional.Bool) -  Whether it should be a identification technology
     * @param "LifeSupportTechnology" (optional.Bool) -  Whether it should be a life support technology
     * @param "SensorTechnology" (optional.Bool) -  Whether it should be a sensor technology
     * @param "ShieldTechnology" (optional.Bool) -  Whether it should be a shield technology
     * @param "Tool" (optional.Bool) -  Whether it should be a tool
     * @param "CulinaryTool" (optional.Bool) -  Whether it should be a culinary tool
     * @param "EngineeringTool" (optional.Bool) -  Whether it should be a engineering tool
     * @param "HouseholdTool" (optional.Bool) -  Whether it should be a household tool
     * @param "MedicalEquipment" (optional.Bool) -  Whether it should be a medical equipment
     * @param "TransporterTechnology" (optional.Bool) -  Whether it&#39;s a transporter technology

@return TechnologyBaseResponse
*/

type TechnologySearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Name optional.String
	BorgTechnology optional.Bool
	BorgComponent optional.Bool
	CommunicationsTechnology optional.Bool
	ComputerTechnology optional.Bool
	ComputerProgramming optional.Bool
	Subroutine optional.Bool
	Database optional.Bool
	EnergyTechnology optional.Bool
	FictionalTechnology optional.Bool
	HolographicTechnology optional.Bool
	IdentificationTechnology optional.Bool
	LifeSupportTechnology optional.Bool
	SensorTechnology optional.Bool
	ShieldTechnology optional.Bool
	Tool optional.Bool
	CulinaryTool optional.Bool
	EngineeringTool optional.Bool
	HouseholdTool optional.Bool
	MedicalEquipment optional.Bool
	TransporterTechnology optional.Bool
}

func (a *TechnologyApiService) TechnologySearchPost(ctx context.Context, localVarOptionals *TechnologySearchPostOpts) (TechnologyBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TechnologyBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/technology/search"

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
	if localVarOptionals != nil && localVarOptionals.BorgTechnology.IsSet() {
		localVarFormParams.Add("borgTechnology", parameterToString(localVarOptionals.BorgTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BorgComponent.IsSet() {
		localVarFormParams.Add("borgComponent", parameterToString(localVarOptionals.BorgComponent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommunicationsTechnology.IsSet() {
		localVarFormParams.Add("communicationsTechnology", parameterToString(localVarOptionals.CommunicationsTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ComputerTechnology.IsSet() {
		localVarFormParams.Add("computerTechnology", parameterToString(localVarOptionals.ComputerTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ComputerProgramming.IsSet() {
		localVarFormParams.Add("computerProgramming", parameterToString(localVarOptionals.ComputerProgramming.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Subroutine.IsSet() {
		localVarFormParams.Add("subroutine", parameterToString(localVarOptionals.Subroutine.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Database.IsSet() {
		localVarFormParams.Add("database", parameterToString(localVarOptionals.Database.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnergyTechnology.IsSet() {
		localVarFormParams.Add("energyTechnology", parameterToString(localVarOptionals.EnergyTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FictionalTechnology.IsSet() {
		localVarFormParams.Add("fictionalTechnology", parameterToString(localVarOptionals.FictionalTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HolographicTechnology.IsSet() {
		localVarFormParams.Add("holographicTechnology", parameterToString(localVarOptionals.HolographicTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdentificationTechnology.IsSet() {
		localVarFormParams.Add("identificationTechnology", parameterToString(localVarOptionals.IdentificationTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LifeSupportTechnology.IsSet() {
		localVarFormParams.Add("lifeSupportTechnology", parameterToString(localVarOptionals.LifeSupportTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SensorTechnology.IsSet() {
		localVarFormParams.Add("sensorTechnology", parameterToString(localVarOptionals.SensorTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ShieldTechnology.IsSet() {
		localVarFormParams.Add("shieldTechnology", parameterToString(localVarOptionals.ShieldTechnology.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tool.IsSet() {
		localVarFormParams.Add("tool", parameterToString(localVarOptionals.Tool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CulinaryTool.IsSet() {
		localVarFormParams.Add("culinaryTool", parameterToString(localVarOptionals.CulinaryTool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EngineeringTool.IsSet() {
		localVarFormParams.Add("engineeringTool", parameterToString(localVarOptionals.EngineeringTool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HouseholdTool.IsSet() {
		localVarFormParams.Add("householdTool", parameterToString(localVarOptionals.HouseholdTool.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MedicalEquipment.IsSet() {
		localVarFormParams.Add("medicalEquipment", parameterToString(localVarOptionals.MedicalEquipment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransporterTechnology.IsSet() {
		localVarFormParams.Add("transporterTechnology", parameterToString(localVarOptionals.TransporterTechnology.Value(), ""))
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
			var v TechnologyBaseResponse
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
