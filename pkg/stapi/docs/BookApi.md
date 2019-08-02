# \BookApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookGet**](BookApi.md#BookGet) | **Get** /book | 
[**BookSearchGet**](BookApi.md#BookSearchGet) | **Get** /book/search | 
[**BookSearchPost**](BookApi.md#BookSearchPost) | **Post** /book/search | 


# **BookGet**
> BookFullResponse BookGet(ctx, uid, optional)


Retrival of a single book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Book unique ID | 
 **optional** | ***BookGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**BookFullResponse**](BookFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookSearchGet**
> BookBaseResponse BookSearchGet(ctx, optional)


Pagination over books

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**BookBaseResponse**](BookBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookSearchPost**
> BookBaseResponse BookSearchPost(ctx, optional)


Searching books

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Book title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the book was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the book was published | 
 **numberOfPagesFrom** | **optional.Int32**| Minimal number of pages | 
 **numberOfPagesTo** | **optional.Int32**| Maximal number of pages | 
 **stardateFrom** | **optional.Float32**| Starting stardate of book story | 
 **stardateTo** | **optional.Float32**| Ending stardate of book story | 
 **yearFrom** | **optional.Int32**| Starting year of book story | 
 **yearTo** | **optional.Int32**| Ending year of book story | 
 **novel** | **optional.Bool**| Whether it should be a novel | 
 **referenceBook** | **optional.Bool**| Whether it should be a reference book | 
 **biographyBook** | **optional.Bool**| Whether it should be a biography book | 
 **rolePlayingBook** | **optional.Bool**| Whether it should be a role playing book | 
 **eBook** | **optional.Bool**| Whether it should be an e-book | 
 **anthology** | **optional.Bool**| Whether it should be an anthology | 
 **novelization** | **optional.Bool**| Whether it should be novelization | 
 **audiobook** | **optional.Bool**| Whether it should be an audiobook | 
 **audiobookAbridged** | **optional.Bool**| Whether it should be an audiobook, abridged | 
 **audiobookPublishedYearFrom** | **optional.Int32**| Starting year the audiobook was published | 
 **audiobookPublishedYearTo** | **optional.Int32**| Ending year the audiobook was published | 
 **audiobookRunTimeFrom** | **optional.Int32**| Minimal audiobook run time, in minutes | 
 **audiobookRunTimeTo** | **optional.Int32**| Maximal audiobook run time, in minutes | 

### Return type

[**BookBaseResponse**](BookBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

