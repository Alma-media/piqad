# \SpacecraftClassApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpacecraftClassGet**](SpacecraftClassApi.md#SpacecraftClassGet) | **Get** /spacecraftClass | 
[**SpacecraftClassSearchGet**](SpacecraftClassApi.md#SpacecraftClassSearchGet) | **Get** /spacecraftClass/search | 
[**SpacecraftClassSearchPost**](SpacecraftClassApi.md#SpacecraftClassSearchPost) | **Post** /spacecraftClass/search | 


# **SpacecraftClassGet**
> SpacecraftClassFullResponse SpacecraftClassGet(ctx, uid, optional)


Retrival of a single spacecraft class

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| SpacecraftClass unique ID | 
 **optional** | ***SpacecraftClassGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacecraftClassGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**SpacecraftClassFullResponse**](SpacecraftClassFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpacecraftClassSearchGet**
> SpacecraftClassBaseResponse SpacecraftClassSearchGet(ctx, optional)


Pagination over spacecraft classes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpacecraftClassSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacecraftClassSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**SpacecraftClassBaseResponse**](SpacecraftClassBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SpacecraftClassSearchPost**
> SpacecraftClassBaseResponse SpacecraftClassSearchPost(ctx, optional)


Searching spacecraft classes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpacecraftClassSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpacecraftClassSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Spacecraft class name | 
 **warpCapableSpecies** | **optional.Bool**| Whether it should be a warp-capable spacecraft class | 
 **alternateReality** | **optional.Bool**| Whether this spacecraft class should be from alternate reality | 

### Return type

[**SpacecraftClassBaseResponse**](SpacecraftClassBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

