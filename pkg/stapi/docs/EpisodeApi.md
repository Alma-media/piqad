# \EpisodeApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpisodeGet**](EpisodeApi.md#EpisodeGet) | **Get** /episode | 
[**EpisodeSearchGet**](EpisodeApi.md#EpisodeSearchGet) | **Get** /episode/search | 
[**EpisodeSearchPost**](EpisodeApi.md#EpisodeSearchPost) | **Post** /episode/search | 


# **EpisodeGet**
> EpisodeFullResponse EpisodeGet(ctx, uid, optional)


Retrival of a single episode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Episode unique ID | 
 **optional** | ***EpisodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EpisodeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**EpisodeFullResponse**](EpisodeFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EpisodeSearchGet**
> EpisodeBaseResponse EpisodeSearchGet(ctx, optional)


Pagination over episodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EpisodeSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EpisodeSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**EpisodeBaseResponse**](EpisodeBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EpisodeSearchPost**
> EpisodeBaseResponse EpisodeSearchPost(ctx, optional)


Searching episodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EpisodeSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EpisodeSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Episode title | 
 **seasonNumberFrom** | **optional.Int32**| Minimal season number | 
 **seasonNumberTo** | **optional.Int32**| Maximal season number | 
 **episodeNumberFrom** | **optional.Int32**| Minimal episode number in season | 
 **episodeNumberTo** | **optional.Int32**| Maximal episode number in season | 
 **productionSerialNumber** | **optional.String**| Production serial number | 
 **featureLength** | **optional.Bool**| Whether it should be a feature length episode | 
 **stardateFrom** | **optional.Float32**| Starting stardate of episode story | 
 **stardateTo** | **optional.Float32**| Ending stardate of episode story | 
 **yearFrom** | **optional.Int32**| Starting year of episode story | 
 **yearTo** | **optional.Int32**| Ending year of episode story | 
 **usAirDateFrom** | **optional.String**| Minimal date the episode was first aired in the United States | 
 **usAirDateTo** | **optional.String**| Maximal date the episode was first aired in the United States | 
 **finalScriptDateFrom** | **optional.String**| Minimal date the episode script was completed | 
 **finalScriptDateTo** | **optional.String**| Maximal date the episode script was completed | 

### Return type

[**EpisodeBaseResponse**](EpisodeBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

