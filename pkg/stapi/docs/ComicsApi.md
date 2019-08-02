# \ComicsApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComicsGet**](ComicsApi.md#ComicsGet) | **Get** /comics | 
[**ComicsSearchGet**](ComicsApi.md#ComicsSearchGet) | **Get** /comics/search | 
[**ComicsSearchPost**](ComicsApi.md#ComicsSearchPost) | **Post** /comics/search | 


# **ComicsGet**
> ComicsFullResponse ComicsGet(ctx, uid, optional)


Retrival of a single comics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Comics unique ID | 
 **optional** | ***ComicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicsFullResponse**](ComicsFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicsSearchGet**
> ComicsBaseResponse ComicsSearchGet(ctx, optional)


Pagination over comics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicsSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicsSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ComicsBaseResponse**](ComicsBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComicsSearchPost**
> ComicsBaseResponse ComicsSearchPost(ctx, optional)


Searching comics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComicsSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComicsSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Comics title | 
 **publishedYearFrom** | **optional.Int32**| Starting year the comics was published | 
 **publishedYearTo** | **optional.Int32**| Ending year the comics was published | 
 **numberOfPagesFrom** | **optional.Int32**| Minimal number of pages | 
 **numberOfPagesTo** | **optional.Int32**| Maximal number of pages | 
 **stardateFrom** | **optional.Float32**| Starting stardate of comics story | 
 **stardateTo** | **optional.Float32**| Ending stardate of comics story | 
 **yearFrom** | **optional.Int32**| Starting year of comics story | 
 **yearTo** | **optional.Int32**| Ending year of comics story | 
 **photonovel** | **optional.Bool**| Whether it should be a photonovel | 
 **adaptation** | **optional.Bool**| Whether it should be an adaptation of an episode or a movie | 

### Return type

[**ComicsBaseResponse**](ComicsBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

