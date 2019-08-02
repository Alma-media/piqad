# \BookCollectionApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookCollectionGet**](BookCollectionApi.md#BookCollectionGet) | **Get** /bookCollection | 
[**BookCollectionSearchGet**](BookCollectionApi.md#BookCollectionSearchGet) | **Get** /bookCollection/search | 
[**BookCollectionSearchPost**](BookCollectionApi.md#BookCollectionSearchPost) | **Post** /bookCollection/search | 


# **BookCollectionGet**
> BookCollectionFullResponse BookCollectionGet(ctx, uid, optional)


Retrival of a single book collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Book collection unique ID | 
 **optional** | ***BookCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**BookCollectionFullResponse**](BookCollectionFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookCollectionSearchGet**
> BookCollectionBaseResponse BookCollectionSearchGet(ctx, optional)


Pagination over book collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookCollectionSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookCollectionSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**BookCollectionBaseResponse**](BookCollectionBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookCollectionSearchPost**
> BookCollectionBaseResponse BookCollectionSearchPost(ctx, optional)


Searching book collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookCollectionSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookCollectionSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Book collection title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the book collection was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the book collection was published | 
 **numberOfPagesFrom** | **optional.Int32**| Minimal number of pages | 
 **numberOfPagesTo** | **optional.Int32**| Maximal number of pages | 
 **stardateFrom** | **optional.Float32**| Starting stardate of book collection stories | 
 **stardateTo** | **optional.Float32**| Ending stardate of book collections stories | 
 **yearFrom** | **optional.Int32**| Starting year of book collection stories | 
 **yearTo** | **optional.Int32**| Ending year of book collections stories | 

### Return type

[**BookCollectionBaseResponse**](BookCollectionBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

