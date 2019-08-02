# \WeaponApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WeaponGet**](WeaponApi.md#WeaponGet) | **Get** /weapon | 
[**WeaponSearchGet**](WeaponApi.md#WeaponSearchGet) | **Get** /weapon/search | 
[**WeaponSearchPost**](WeaponApi.md#WeaponSearchPost) | **Post** /weapon/search | 


# **WeaponGet**
> WeaponFullResponse WeaponGet(ctx, uid, optional)


Retrival of a single weapon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Weapon unique ID | 
 **optional** | ***WeaponGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WeaponGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**WeaponFullResponse**](WeaponFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WeaponSearchGet**
> WeaponBaseResponse WeaponSearchGet(ctx, optional)


Pagination over weapons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WeaponSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WeaponSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**WeaponBaseResponse**](WeaponBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WeaponSearchPost**
> WeaponBaseResponse WeaponSearchPost(ctx, optional)


Searching weapons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WeaponSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WeaponSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Weapon name | 
 **handHeldWeapon** | **optional.Bool**| Whether it should be a hand-help weapon | 
 **laserTechnology** | **optional.Bool**| Whether it should be a laser technology | 
 **plasmaTechnology** | **optional.Bool**| Whether it should be a plasma technology | 
 **photonicTechnology** | **optional.Bool**| Whether it should be a photonic technology | 
 **phaserTechnology** | **optional.Bool**| Whether it should be a phaser technology | 
 **mirror** | **optional.Bool**| Whether this weapon should be from mirror universe | 
 **alternateReality** | **optional.Bool**| Whether this weapon should be from alternate reality | 

### Return type

[**WeaponBaseResponse**](WeaponBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

