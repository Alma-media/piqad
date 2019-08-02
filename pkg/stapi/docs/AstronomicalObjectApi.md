# \AstronomicalObjectApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AstronomicalObjectGet**](AstronomicalObjectApi.md#AstronomicalObjectGet) | **Get** /astronomicalObject | 
[**AstronomicalObjectSearchGet**](AstronomicalObjectApi.md#AstronomicalObjectSearchGet) | **Get** /astronomicalObject/search | 
[**AstronomicalObjectSearchPost**](AstronomicalObjectApi.md#AstronomicalObjectSearchPost) | **Post** /astronomicalObject/search | 


# **AstronomicalObjectGet**
> AstronomicalObjectFullResponse AstronomicalObjectGet(ctx, uid, optional)


Retrival of a single astronomical object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Astronomical object&#39;s unique ID | 
 **optional** | ***AstronomicalObjectGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AstronomicalObjectGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**AstronomicalObjectFullResponse**](AstronomicalObjectFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AstronomicalObjectSearchGet**
> AstronomicalObjectBaseResponse AstronomicalObjectSearchGet(ctx, optional)


Pagination over astronomical objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AstronomicalObjectSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AstronomicalObjectSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**AstronomicalObjectBaseResponse**](AstronomicalObjectBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AstronomicalObjectSearchPost**
> AstronomicalObjectBaseResponse AstronomicalObjectSearchPost(ctx, optional)


Searching astronomical objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AstronomicalObjectSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AstronomicalObjectSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Astronomical object name | 
 **astronomicalObjectType** | **optional.String**| Type of astronomical object | 
 **locationUid** | **optional.String**| Unique ID of astronomical object containing objects being searched | 

### Return type

[**AstronomicalObjectBaseResponse**](AstronomicalObjectBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

