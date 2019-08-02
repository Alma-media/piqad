# \SpeciesApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpeciesGet**](SpeciesApi.md#SpeciesGet) | **Get** /species | 
[**SpeciesSearchGet**](SpeciesApi.md#SpeciesSearchGet) | **Get** /species/search | 
[**SpeciesSearchPost**](SpeciesApi.md#SpeciesSearchPost) | **Post** /species/search | 


# **SpeciesGet**
> SpeciesFullResponse SpeciesGet(ctx, uid, optional)


Retrival of a single species

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Species unique ID | 
 **optional** | ***SpeciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpeciesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**SpeciesFullResponse**](SpeciesFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpeciesSearchGet**
> SpeciesBaseResponse SpeciesSearchGet(ctx, optional)


Pagination over species

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpeciesSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpeciesSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**SpeciesBaseResponse**](SpeciesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpeciesSearchPost**
> SpeciesBaseResponse SpeciesSearchPost(ctx, optional)


Searching species

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpeciesSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpeciesSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Species name | 
 **extinctSpecies** | **optional.Bool**| Whether it should be an extinct species | 
 **warpCapableSpecies** | **optional.Bool**| Whether it should be a warp-capable species | 
 **extraGalacticSpecies** | **optional.Bool**| Whether it should be an extra-galactic species | 
 **humanoidSpecies** | **optional.Bool**| Whether it should be a humanoid species | 
 **reptilianSpecies** | **optional.Bool**| Whether it should be a reptilian species | 
 **nonCorporealSpecies** | **optional.Bool**| Whether it should be a non-corporeal species | 
 **shapeshiftingSpecies** | **optional.Bool**| Whether it should be a shapeshifting species | 
 **spaceborneSpecies** | **optional.Bool**| Whether it should be a spaceborne species | 
 **telepathicSpecies** | **optional.Bool**| Whether it should be a telepathic species | 
 **transDimensionalSpecies** | **optional.Bool**| Whether it should be a trans-dimensional species | 
 **unnamedSpecies** | **optional.Bool**| Whether it should be a unnamed species | 
 **alternateReality** | **optional.Bool**| Whether this species should be from alternate reality | 

### Return type

[**SpeciesBaseResponse**](SpeciesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

