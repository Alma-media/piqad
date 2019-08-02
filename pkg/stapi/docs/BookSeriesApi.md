# \BookSeriesApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookSeriesGet**](BookSeriesApi.md#BookSeriesGet) | **Get** /bookSeries | 
[**BookSeriesSearchGet**](BookSeriesApi.md#BookSeriesSearchGet) | **Get** /bookSeries/search | 
[**BookSeriesSearchPost**](BookSeriesApi.md#BookSeriesSearchPost) | **Post** /bookSeries/search | 


# **BookSeriesGet**
> BookSeriesFullResponse BookSeriesGet(ctx, uid, optional)


Retrival of a single book series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Book series unique ID | 
 **optional** | ***BookSeriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookSeriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**BookSeriesFullResponse**](BookSeriesFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookSeriesSearchGet**
> BookSeriesBaseResponse BookSeriesSearchGet(ctx, optional)


Pagination over book series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookSeriesSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookSeriesSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**BookSeriesBaseResponse**](BookSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookSeriesSearchPost**
> BookSeriesBaseResponse BookSeriesSearchPost(ctx, optional)


Searching book series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookSeriesSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookSeriesSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Book series title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the book series was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the book series was published | 
 **numberOfBooksFrom** | **optional.Int32**| Minimal number of books | 
 **numberOfBooksTo** | **optional.Int32**| Maximal number of books | 
 **yearFrom** | **optional.Int32**| Starting year of book series stories | 
 **yearTo** | **optional.Int32**| Ending year of book series stories | 
 **miniseries** | **optional.Bool**| Whether it should be a miniseries | 
 **eBookSeries** | **optional.Bool**| Whether it should be an e-book series | 

### Return type

[**BookSeriesBaseResponse**](BookSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

