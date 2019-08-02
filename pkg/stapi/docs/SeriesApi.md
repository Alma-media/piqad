# \SeriesApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeriesGet**](SeriesApi.md#SeriesGet) | **Get** /series | 
[**SeriesSearchGet**](SeriesApi.md#SeriesSearchGet) | **Get** /series/search | 
[**SeriesSearchPost**](SeriesApi.md#SeriesSearchPost) | **Post** /series/search | 


# **SeriesGet**
> SeriesFullResponse SeriesGet(ctx, uid, optional)


Retrival of a single series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Series unique ID | 
 **optional** | ***SeriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**SeriesFullResponse**](SeriesFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SeriesSearchGet**
> SeriesBaseResponse SeriesSearchGet(ctx, optional)


Pagination over series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SeriesSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**SeriesBaseResponse**](SeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SeriesSearchPost**
> SeriesBaseResponse SeriesSearchPost(ctx, optional)


Searching series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SeriesSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Series title | 
 **abbreviation** | **optional.String**| Series abbreviation | 
 **productionStartYearFrom** | **optional.Int32**| Minimal year the series production started | 
 **productionStartYearTo** | **optional.Int32**| Maximal year the series production started | 
 **productionEndYearFrom** | **optional.Int32**| Minimal year the series production ended | 
 **productionEndYearTo** | **optional.Int32**| Maximal year the series production ended | 
 **originalRunStartDateFrom** | **optional.String**| Minimal date the series originally ran from | 
 **originalRunStartDateTo** | **optional.String**| Maximal date the series originally ran from | 
 **originalRunEndDateFrom** | **optional.String**| Minimal date the series originally ran to | 
 **originalRunEndDateTo** | **optional.String**| Maximal date the series originally ran to | 

### Return type

[**SeriesBaseResponse**](SeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

