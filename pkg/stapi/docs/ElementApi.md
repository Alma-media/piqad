# \ElementApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElementGet**](ElementApi.md#ElementGet) | **Get** /element | 
[**ElementSearchGet**](ElementApi.md#ElementSearchGet) | **Get** /element/search | 
[**ElementSearchPost**](ElementApi.md#ElementSearchPost) | **Post** /element/search | 


# **ElementGet**
> ElementFullResponse ElementGet(ctx, uid, optional)


Retrival of a single element

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Element unique ID | 
 **optional** | ***ElementGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ElementFullResponse**](ElementFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementSearchGet**
> ElementBaseResponse ElementSearchGet(ctx, optional)


Pagination over elements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ElementSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ElementBaseResponse**](ElementBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementSearchPost**
> ElementBaseResponse ElementSearchPost(ctx, optional)


Searching elements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ElementSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Element name | 
 **symbol** | **optional.String**| Element symbol | 
 **transuranium** | **optional.Bool**| Whether it should be a transuranium | 
 **gammaSeries** | **optional.Bool**| Whether it should belong to Gamma series | 
 **hypersonicSeries** | **optional.Bool**| Whether it should belong to Hypersonic series | 
 **megaSeries** | **optional.Bool**| Whether it should belong to Mega series | 
 **omegaSeries** | **optional.Bool**| Whether it should belong to Omega series | 
 **transonicSeries** | **optional.Bool**| Whether it should belong to Transonic series | 
 **worldSeries** | **optional.Bool**| Whether it should belong to World series | 

### Return type

[**ElementBaseResponse**](ElementBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

