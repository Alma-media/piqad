# \SoundtrackApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoundtrackGet**](SoundtrackApi.md#SoundtrackGet) | **Get** /soundtrack | 
[**SoundtrackSearchGet**](SoundtrackApi.md#SoundtrackSearchGet) | **Get** /soundtrack/search | 
[**SoundtrackSearchPost**](SoundtrackApi.md#SoundtrackSearchPost) | **Post** /soundtrack/search | 


# **SoundtrackGet**
> SoundtrackFullResponse SoundtrackGet(ctx, uid, optional)


Retrival of a single soundtrack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Soundtrack unique ID | 
 **optional** | ***SoundtrackGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoundtrackGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**SoundtrackFullResponse**](SoundtrackFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoundtrackSearchGet**
> SoundtrackBaseResponse SoundtrackSearchGet(ctx, optional)


Pagination over soundtracks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundtrackSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoundtrackSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**SoundtrackBaseResponse**](SoundtrackBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoundtrackSearchPost**
> SoundtrackBaseResponse SoundtrackSearchPost(ctx, optional)


Searching soundtracks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundtrackSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoundtrackSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Soundtrack title | 
 **releaseDateFrom** | **optional.String**| Minimal release date | 
 **releaseDateTo** | **optional.String**| Maximal release date | 
 **lengthFrom** | **optional.Int32**| Minimal length, in seconds | 
 **lengthTo** | **optional.Int32**| Maximal length, in seconds | 

### Return type

[**SoundtrackBaseResponse**](SoundtrackBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

