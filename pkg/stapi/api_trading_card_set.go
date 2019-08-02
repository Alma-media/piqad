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

type TradingCardSetApiService service

/* 
TradingCardSetApiService
Retrival of a single trading card set
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uid Trading card set unique ID
 * @param optional nil or *TradingCardSetGetOpts - Optional Parameters:
     * @param "ApiKey" (optional.String) -  API key

@return TradingCardSetFullResponse
*/

type TradingCardSetGetOpts struct { 
	ApiKey optional.String
}

func (a *TradingCardSetApiService) TradingCardSetGet(ctx context.Context, uid string, localVarOptionals *TradingCardSetGetOpts) (TradingCardSetFullResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TradingCardSetFullResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tradingCardSet"

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
			var v TradingCardSetFullResponse
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
TradingCardSetApiService
Pagination over trading card sets
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TradingCardSetSearchGetOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "ApiKey" (optional.String) -  API key

@return TradingCardSetBaseResponse
*/

type TradingCardSetSearchGetOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	ApiKey optional.String
}

func (a *TradingCardSetApiService) TradingCardSetSearchGet(ctx context.Context, localVarOptionals *TradingCardSetSearchGetOpts) (TradingCardSetBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TradingCardSetBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tradingCardSet/search"

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
			var v TradingCardSetBaseResponse
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
TradingCardSetApiService
Searching trading card sets
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TradingCardSetSearchPostOpts - Optional Parameters:
     * @param "PageNumber" (optional.Int32) -  Zero-based page number
     * @param "PageSize" (optional.Int32) -  Page size
     * @param "Sort" (optional.String) -  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC
     * @param "ApiKey" (optional.String) -  API key
     * @param "Name" (optional.String) -  Trading card set name
     * @param "ReleaseYearFrom" (optional.Int32) -  Starting year the trading card set was released
     * @param "ReleaseYearTo" (optional.Int32) -  Ending year the trading card set was released
     * @param "CardsPerPackFrom" (optional.Int32) -  Minimal number of cards per deck
     * @param "CardsPerPackTo" (optional.Int32) -  Minimal number of cards per deck
     * @param "PacksPerBoxFrom" (optional.Int32) -  Minimal number of packs per box
     * @param "PacksPerBoxTo" (optional.Int32) -  Minimal number of packs per box
     * @param "BoxesPerCaseFrom" (optional.Int32) -  Minimal number of boxes per case
     * @param "BoxesPerCaseTo" (optional.Int32) -  Minimal number of boxes per case
     * @param "ProductionRunFrom" (optional.Int32) -  Minimal production run
     * @param "ProductionRunTo" (optional.Int32) -  Minimal production run
     * @param "ProductionRunUnit" (optional.String) -  Production run unit
     * @param "CardWidthFrom" (optional.Float64) -  Minimal card width, in inches
     * @param "CardWidthTo" (optional.Float64) -  Minimal card width, in inches
     * @param "CardHeightFrom" (optional.Float64) -  Minimal card height, in inches
     * @param "CardHeightTo" (optional.Float64) -  Minimal card height, in inches

@return TradingCardSetBaseResponse
*/

type TradingCardSetSearchPostOpts struct { 
	PageNumber optional.Int32
	PageSize optional.Int32
	Sort optional.String
	ApiKey optional.String
	Name optional.String
	ReleaseYearFrom optional.Int32
	ReleaseYearTo optional.Int32
	CardsPerPackFrom optional.Int32
	CardsPerPackTo optional.Int32
	PacksPerBoxFrom optional.Int32
	PacksPerBoxTo optional.Int32
	BoxesPerCaseFrom optional.Int32
	BoxesPerCaseTo optional.Int32
	ProductionRunFrom optional.Int32
	ProductionRunTo optional.Int32
	ProductionRunUnit optional.String
	CardWidthFrom optional.Float64
	CardWidthTo optional.Float64
	CardHeightFrom optional.Float64
	CardHeightTo optional.Float64
}

func (a *TradingCardSetApiService) TradingCardSetSearchPost(ctx context.Context, localVarOptionals *TradingCardSetSearchPostOpts) (TradingCardSetBaseResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TradingCardSetBaseResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tradingCardSet/search"

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
	if localVarOptionals != nil && localVarOptionals.ReleaseYearFrom.IsSet() {
		localVarFormParams.Add("releaseYearFrom", parameterToString(localVarOptionals.ReleaseYearFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReleaseYearTo.IsSet() {
		localVarFormParams.Add("releaseYearTo", parameterToString(localVarOptionals.ReleaseYearTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardsPerPackFrom.IsSet() {
		localVarFormParams.Add("cardsPerPackFrom", parameterToString(localVarOptionals.CardsPerPackFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardsPerPackTo.IsSet() {
		localVarFormParams.Add("cardsPerPackTo", parameterToString(localVarOptionals.CardsPerPackTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PacksPerBoxFrom.IsSet() {
		localVarFormParams.Add("packsPerBoxFrom", parameterToString(localVarOptionals.PacksPerBoxFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PacksPerBoxTo.IsSet() {
		localVarFormParams.Add("packsPerBoxTo", parameterToString(localVarOptionals.PacksPerBoxTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BoxesPerCaseFrom.IsSet() {
		localVarFormParams.Add("boxesPerCaseFrom", parameterToString(localVarOptionals.BoxesPerCaseFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BoxesPerCaseTo.IsSet() {
		localVarFormParams.Add("boxesPerCaseTo", parameterToString(localVarOptionals.BoxesPerCaseTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductionRunFrom.IsSet() {
		localVarFormParams.Add("productionRunFrom", parameterToString(localVarOptionals.ProductionRunFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductionRunTo.IsSet() {
		localVarFormParams.Add("productionRunTo", parameterToString(localVarOptionals.ProductionRunTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductionRunUnit.IsSet() {
		localVarFormParams.Add("productionRunUnit", parameterToString(localVarOptionals.ProductionRunUnit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardWidthFrom.IsSet() {
		localVarFormParams.Add("cardWidthFrom", parameterToString(localVarOptionals.CardWidthFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardWidthTo.IsSet() {
		localVarFormParams.Add("cardWidthTo", parameterToString(localVarOptionals.CardWidthTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardHeightFrom.IsSet() {
		localVarFormParams.Add("cardHeightFrom", parameterToString(localVarOptionals.CardHeightFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardHeightTo.IsSet() {
		localVarFormParams.Add("cardHeightTo", parameterToString(localVarOptionals.CardHeightTo.Value(), ""))
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
			var v TradingCardSetBaseResponse
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
