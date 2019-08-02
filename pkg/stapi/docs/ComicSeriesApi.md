# \ComicSeriesApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComicSeriesGet**](ComicSeriesApi.md#ComicSeriesGet) | **Get** /comicSeries | 
[**ComicSeriesSearchGet**](ComicSeriesApi.md#ComicSeriesSearchGet) | **Get** /comicSeries/search | 
[**ComicSeriesSearchPost**](ComicSeriesApi.md#ComicSeriesSearchPost) | **Post** /comicSeries/search | 


# **ComicSeriesGet**
> ComicSeriesFullResponse ComicSeriesGet(ctx, uid, optional)


Retrival of a single comic series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Comic series unique ID | 
 **optional** | ***ComicSeriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicSeriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicSeriesFullResponse**](ComicSeriesFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicSeriesSearchGet**
> ComicSeriesBaseResponse ComicSeriesSearchGet(ctx, optional)


Pagination over comic series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicSeriesSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicSeriesSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicSeriesBaseResponse**](ComicSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicSeriesSearchPost**
> ComicSeriesBaseResponse ComicSeriesSearchPost(ctx, optional)


Searching comic series

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicSeriesSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicSeriesSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Comic series title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the comic series was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the comic series was published | 
 **numberOfIssuesFrom** | **optional.Int32**| Minimal number of issues | 
 **numberOfIssuesTo** | **optional.Int32**| Maximal number of issues | 
 **stardateFrom** | **optional.Float32**| Starting stardate of comic series stories | 
 **stardateTo** | **optional.Float32**| Starting stardate of comic series stories | 
 **yearFrom** | **optional.Int32**| Starting year of comic series stories | 
 **yearTo** | **optional.Int32**| Ending year of comic series stories | 
 **miniseries** | **optional.Bool**| Whether it should be a miniseries | 
 **photonovelSeries** | **optional.Bool**| Whether it should be photonovel series | 

### Return type

[**ComicSeriesBaseResponse**](ComicSeriesBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

