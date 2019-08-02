# \ComicCollectionApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComicCollectionGet**](ComicCollectionApi.md#ComicCollectionGet) | **Get** /comicCollection | 
[**ComicCollectionSearchGet**](ComicCollectionApi.md#ComicCollectionSearchGet) | **Get** /comicCollection/search | 
[**ComicCollectionSearchPost**](ComicCollectionApi.md#ComicCollectionSearchPost) | **Post** /comicCollection/search | 


# **ComicCollectionGet**
> ComicCollectionFullResponse ComicCollectionGet(ctx, uid, optional)


Retrival of a single comic collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Comic collection unique ID | 
 **optional** | ***ComicCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicCollectionFullResponse**](ComicCollectionFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicCollectionSearchGet**
> ComicCollectionBaseResponse ComicCollectionSearchGet(ctx, optional)


Pagination over comic collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicCollectionSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicCollectionSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicCollectionBaseResponse**](ComicCollectionBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicCollectionSearchPost**
> ComicCollectionBaseResponse ComicCollectionSearchPost(ctx, optional)


Searching comic collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicCollectionSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicCollectionSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Comic collection title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the comic collection was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the comic collection was published | 
 **numberOfPagesFrom** | **optional.Int32**| Minimal number of pages | 
 **numberOfPagesTo** | **optional.Int32**| Maximal number of pages | 
 **stardateFrom** | **optional.Float32**| Starting stardate of comic collection stories | 
 **stardateTo** | **optional.Float32**| Ending stardate of comic collections stories | 
 **yearFrom** | **optional.Int32**| Starting year of comic collection stories | 
 **yearTo** | **optional.Int32**| Ending year of comic collections stories | 
 **photonovel** | **optional.Bool**| Whether it should be an photonovel collection | 

### Return type

[**ComicCollectionBaseResponse**](ComicCollectionBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

