# \CharacterApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CharacterGet**](CharacterApi.md#CharacterGet) | **Get** /character | 
[**CharacterSearchGet**](CharacterApi.md#CharacterSearchGet) | **Get** /character/search | 
[**CharacterSearchPost**](CharacterApi.md#CharacterSearchPost) | **Post** /character/search | 


# **CharacterGet**
> CharacterFullResponse CharacterGet(ctx, uid, optional)


Retrival of a single character

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Character unique ID | 
 **optional** | ***CharacterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**CharacterFullResponse**](CharacterFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CharacterSearchGet**
> CharacterBaseResponse CharacterSearchGet(ctx, optional)


Pagination over characters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CharacterSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**CharacterBaseResponse**](CharacterBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CharacterSearchPost**
> CharacterBaseResponse CharacterSearchPost(ctx, optional)


Searching characters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CharacterSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Character name | 
 **gender** | **optional.String**| Character gender | 
 **deceased** | **optional.Bool**| Whether it should be a deceased character | 
 **hologram** | **optional.Bool**| Whether it should be a hologram | 
 **fictionalCharacter** | **optional.Bool**| Whether it should be a fictional character (from universe point of view) | 
 **mirror** | **optional.Bool**| Whether it should be a mirror universe character | 
 **alternateReality** | **optional.Bool**| Whether it should be a alternate reality character | 

### Return type

[**CharacterBaseResponse**](CharacterBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

