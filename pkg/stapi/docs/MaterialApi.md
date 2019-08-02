# \MaterialApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialGet**](MaterialApi.md#MaterialGet) | **Get** /material | 
[**MaterialSearchGet**](MaterialApi.md#MaterialSearchGet) | **Get** /material/search | 
[**MaterialSearchPost**](MaterialApi.md#MaterialSearchPost) | **Post** /material/search | 


# **MaterialGet**
> MaterialFullResponse MaterialGet(ctx, uid, optional)


Retrival of a single material

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Material unique ID | 
 **optional** | ***MaterialGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaterialGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**MaterialFullResponse**](MaterialFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialSearchGet**
> MaterialBaseResponse MaterialSearchGet(ctx, optional)


Pagination over materials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MaterialSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaterialSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**MaterialBaseResponse**](MaterialBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialSearchPost**
> MaterialBaseResponse MaterialSearchPost(ctx, optional)


Searching materials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MaterialSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaterialSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Material name | 
 **chemicalCompound** | **optional.Bool**| Whether it should be a chemical compound | 
 **biochemicalCompound** | **optional.Bool**| Whether it should be a biochemical compound | 
 **drug** | **optional.Bool**| Whether it should be a drug | 
 **poisonousSubstance** | **optional.Bool**| Whether it should be a poisonous substance | 
 **explosive** | **optional.Bool**| Whether it should be an explosive | 
 **gemstone** | **optional.Bool**| Whether it should be a gemstone | 
 **alloyOrComposite** | **optional.Bool**| Whether it should be an alloy or a composite | 
 **fuel** | **optional.Bool**| Whether it should be a fuel | 
 **mineral** | **optional.Bool**| Whether it should be a mineral | 
 **preciousMaterial** | **optional.Bool**| Whether it should be a precious material | 

### Return type

[**MaterialBaseResponse**](MaterialBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

