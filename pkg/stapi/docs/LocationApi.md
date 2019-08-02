# \LocationApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationGet**](LocationApi.md#LocationGet) | **Get** /location | 
[**LocationSearchGet**](LocationApi.md#LocationSearchGet) | **Get** /location/search | 
[**LocationSearchPost**](LocationApi.md#LocationSearchPost) | **Post** /location/search | 


# **LocationGet**
> LocationFullResponse LocationGet(ctx, uid, optional)


Retrival of a single location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Location unique ID | 
 **optional** | ***LocationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**LocationFullResponse**](LocationFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationSearchGet**
> LocationBaseResponse LocationSearchGet(ctx, optional)


Pagination over locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**LocationBaseResponse**](LocationBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationSearchPost**
> LocationBaseResponse LocationSearchPost(ctx, optional)


Searching locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Location name | 
 **earthlyLocation** | **optional.Bool**| Whether it should be an earthly location | 
 **fictionalLocation** | **optional.Bool**| Whether it should be a fictional location | 
 **religiousLocation** | **optional.Bool**| Whether it should be a religious location | 
 **geographicalLocation** | **optional.Bool**| Whether it should be a geographical location | 
 **bodyOfWater** | **optional.Bool**| Whether it should be a body of water | 
 **country** | **optional.Bool**| Whether it should be a country | 
 **subnationalEntity** | **optional.Bool**| Whether it should be a subnational entity | 
 **settlement** | **optional.Bool**| Whether it should be a settlement | 
 **usSettlement** | **optional.Bool**| Whether it should be a US settlement | 
 **bajoranSettlement** | **optional.Bool**| Whether it should be a Bajoran settlement | 
 **colony** | **optional.Bool**| Whether it should be a colony | 
 **landform** | **optional.Bool**| Whether it should be a landform | 
 **landmark** | **optional.Bool**| Whether it should be a landmark | 
 **road** | **optional.Bool**| Whether it should be a road | 
 **structure** | **optional.Bool**| Whether it should be a structure | 
 **shipyard** | **optional.Bool**| Whether it should be a shipyard | 
 **buildingInterior** | **optional.Bool**| Whether it should be a building interior | 
 **establishment** | **optional.Bool**| Whether it should be a establishment | 
 **medicalEstablishment** | **optional.Bool**| Whether it should be a medical establishment | 
 **ds9Establishment** | **optional.Bool**| Whether it should be a DS9 establishment | 
 **school** | **optional.Bool**| Whether it should be a school | 
 **mirror** | **optional.Bool**| Whether this location should be from mirror universe | 
 **alternateReality** | **optional.Bool**| Whether this location should be from alternate reality | 

### Return type

[**LocationBaseResponse**](LocationBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

