# \LiteratureApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LiteratureGet**](LiteratureApi.md#LiteratureGet) | **Get** /literature | 
[**LiteratureSearchGet**](LiteratureApi.md#LiteratureSearchGet) | **Get** /literature/search | 
[**LiteratureSearchPost**](LiteratureApi.md#LiteratureSearchPost) | **Post** /literature/search | 


# **LiteratureGet**
> LiteratureFullResponse LiteratureGet(ctx, uid, optional)


Retrival of a single literature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Literature unique ID | 
 **optional** | ***LiteratureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LiteratureGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**LiteratureFullResponse**](LiteratureFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiteratureSearchGet**
> LiteratureBaseResponse LiteratureSearchGet(ctx, optional)


Pagination over literature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LiteratureSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LiteratureSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**LiteratureBaseResponse**](LiteratureBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiteratureSearchPost**
> LiteratureBaseResponse LiteratureSearchPost(ctx, optional)


Searching literature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LiteratureSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LiteratureSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Literature title | 
 **earthlyOrigin** | **optional.Bool**| Whether it should be of earthly origin | 
 **shakespeareanWork** | **optional.Bool**| Whether it should be a Shakespearean work | 
 **report** | **optional.Bool**| Whether it should be a report | 
 **scientificLiterature** | **optional.Bool**| Whether it should be a scientific literature | 
 **technicalManual** | **optional.Bool**| Whether it should be a technical manual | 
 **religiousLiterature** | **optional.Bool**| Whether it should be a religious literature | 

### Return type

[**LiteratureBaseResponse**](LiteratureBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

