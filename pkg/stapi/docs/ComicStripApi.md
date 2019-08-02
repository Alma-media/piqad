# \ComicStripApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComicStripGet**](ComicStripApi.md#ComicStripGet) | **Get** /comicStrip | 
[**ComicStripSearchGet**](ComicStripApi.md#ComicStripSearchGet) | **Get** /comicStrip/search | 
[**ComicStripSearchPost**](ComicStripApi.md#ComicStripSearchPost) | **Post** /comicStrip/search | 


# **ComicStripGet**
> ComicStripFullResponse ComicStripGet(ctx, uid, optional)


Retrival of a single comic strip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Comic strip unique ID | 
 **optional** | ***ComicStripGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicStripGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicStripFullResponse**](ComicStripFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicStripSearchGet**
> ComicStripBaseResponse ComicStripSearchGet(ctx, optional)


Pagination over comic strips

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicStripSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicStripSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicStripBaseResponse**](ComicStripBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicStripSearchPost**
> ComicStripBaseResponse ComicStripSearchPost(ctx, optional)


Searching comic strips

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicStripSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicStripSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Comic strip title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the comic strip was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the comic strip was published | 
 **numberOfPagesFrom** | **optional.Int32**| Minimal number of pages | 
 **numberOfPagesTo** | **optional.Int32**| Maximal number of pages | 
 **yearFrom** | **optional.Int32**| Starting year of comic strip story | 
 **yearTo** | **optional.Int32**| Ending year of comic strip story | 

### Return type

[**ComicStripBaseResponse**](ComicStripBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

