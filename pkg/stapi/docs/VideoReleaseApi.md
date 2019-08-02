# \VideoReleaseApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideoReleaseGet**](VideoReleaseApi.md#VideoReleaseGet) | **Get** /videoRelease | 
[**VideoReleaseSearchGet**](VideoReleaseApi.md#VideoReleaseSearchGet) | **Get** /videoRelease/search | 
[**VideoReleaseSearchPost**](VideoReleaseApi.md#VideoReleaseSearchPost) | **Post** /videoRelease/search | 


# **VideoReleaseGet**
> VideoReleaseFullResponse VideoReleaseGet(ctx, uid, optional)


Retrival of a single video release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Video release unique ID | 
 **optional** | ***VideoReleaseGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideoReleaseGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**VideoReleaseFullResponse**](VideoReleaseFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoReleaseSearchGet**
> VideoReleaseBaseResponse VideoReleaseSearchGet(ctx, optional)


Pagination over video releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VideoReleaseSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideoReleaseSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**VideoReleaseBaseResponse**](VideoReleaseBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoReleaseSearchPost**
> VideoReleaseBaseResponse VideoReleaseSearchPost(ctx, optional)


Searching video releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VideoReleaseSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideoReleaseSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **title** | **optional.String**| Video release title | 
 **yearFrom** | **optional.Int32**| Starting year of video release story | 
 **yearTo** | **optional.Int32**| Ending year of video release story | 
 **runTimeFrom** | **optional.Int32**| Minimal run time, in minutes | 
 **runTimeTo** | **optional.Int32**| Minimal run time, in minutes | 

### Return type

[**VideoReleaseBaseResponse**](VideoReleaseBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

