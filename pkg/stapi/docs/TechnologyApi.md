# \TechnologyApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TechnologyGet**](TechnologyApi.md#TechnologyGet) | **Get** /technology | 
[**TechnologySearchGet**](TechnologyApi.md#TechnologySearchGet) | **Get** /technology/search | 
[**TechnologySearchPost**](TechnologyApi.md#TechnologySearchPost) | **Post** /technology/search | 


# **TechnologyGet**
> TechnologyFullResponse TechnologyGet(ctx, uid, optional)


Retrival of a single technology

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Technology unique ID | 
 **optional** | ***TechnologyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnologyGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**TechnologyFullResponse**](TechnologyFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TechnologySearchGet**
> TechnologyBaseResponse TechnologySearchGet(ctx, optional)


Pagination over technology

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TechnologySearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnologySearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**TechnologyBaseResponse**](TechnologyBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TechnologySearchPost**
> TechnologyBaseResponse TechnologySearchPost(ctx, optional)


Searching technology

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TechnologySearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TechnologySearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Technology name | 
 **borgTechnology** | **optional.Bool**| Whether it should be a Borg technology | 
 **borgComponent** | **optional.Bool**| Whether it should be a Borg component | 
 **communicationsTechnology** | **optional.Bool**| Whether it should be a communications technology | 
 **computerTechnology** | **optional.Bool**| Whether it should be a computer technology | 
 **computerProgramming** | **optional.Bool**| Whether it should be a technology related to computer programming | 
 **subroutine** | **optional.Bool**| Whether it should be a subroutine | 
 **database** | **optional.Bool**| Whether it should be a database | 
 **energyTechnology** | **optional.Bool**| Whether it should be a energy technology | 
 **fictionalTechnology** | **optional.Bool**| Whether it should be a fictional technology | 
 **holographicTechnology** | **optional.Bool**| Whether it should be a holographic technology | 
 **identificationTechnology** | **optional.Bool**| Whether it should be a identification technology | 
 **lifeSupportTechnology** | **optional.Bool**| Whether it should be a life support technology | 
 **sensorTechnology** | **optional.Bool**| Whether it should be a sensor technology | 
 **shieldTechnology** | **optional.Bool**| Whether it should be a shield technology | 
 **tool** | **optional.Bool**| Whether it should be a tool | 
 **culinaryTool** | **optional.Bool**| Whether it should be a culinary tool | 
 **engineeringTool** | **optional.Bool**| Whether it should be a engineering tool | 
 **householdTool** | **optional.Bool**| Whether it should be a household tool | 
 **medicalEquipment** | **optional.Bool**| Whether it should be a medical equipment | 
 **transporterTechnology** | **optional.Bool**| Whether it&#39;s a transporter technology | 

### Return type

[**TechnologyBaseResponse**](TechnologyBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

