# \MagazineSeriesApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MagazineSeriesGet**](MagazineSeriesApi.md#MagazineSeriesGet) | **Get** /magazineSeries | 
[**MagazineSeriesSearchGet**](MagazineSeriesApi.md#MagazineSeriesSearchGet) | **Get** /magazineSeries/search | 
[**MagazineSeriesSearchPost**](MagazineSeriesApi.md#MagazineSeriesSearchPost) | **Post** /magazineSeries/search | 


# **MagazineSeriesGet**
> MagazineSeriesFullResponse MagazineSeriesGet(ctx, uid, optional)


Retrival of a single magazine series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Magazine series unique ID | 
 **optional** | ***MagazineSeriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MagazineSeriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**MagazineSeriesFullResponse**](MagazineSeriesFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MagazineSeriesSearchGet**
> MagazineSeriesBaseResponse MagazineSeriesSearchGet(ctx, optional)


Pagination over magazine series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MagazineSeriesSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MagazineSeriesSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**MagazineSeriesBaseResponse**](MagazineSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MagazineSeriesSearchPost**
> MagazineSeriesBaseResponse MagazineSeriesSearchPost(ctx, optional)


Searching magazine series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MagazineSeriesSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MagazineSeriesSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Magazine series title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the magazine series was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the magazine series was published | 
 **numberOfIssuesFrom** | **optional.Int32**| Minimal number of issues | 
 **numberOfIssuesTo** | **optional.Int32**| Maximal number of issues | 

### Return type

[**MagazineSeriesBaseResponse**](MagazineSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

