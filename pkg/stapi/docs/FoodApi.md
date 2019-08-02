# \FoodApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FoodGet**](FoodApi.md#FoodGet) | **Get** /food | 
[**FoodSearchGet**](FoodApi.md#FoodSearchGet) | **Get** /food/search | 
[**FoodSearchPost**](FoodApi.md#FoodSearchPost) | **Post** /food/search | 


# **FoodGet**
> FoodFullResponse FoodGet(ctx, uid, optional)


Retrival of a single food

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Food unique ID | 
 **optional** | ***FoodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FoodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**FoodFullResponse**](FoodFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FoodSearchGet**
> FoodBaseResponse FoodSearchGet(ctx, optional)


Pagination over foods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FoodSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FoodSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**FoodBaseResponse**](FoodBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FoodSearchPost**
> FoodBaseResponse FoodSearchPost(ctx, optional)


Searching foods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FoodSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FoodSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Food name | 
 **earthlyOrigin** | **optional.Bool**| Whether it should be of earthly origin | 
 **dessert** | **optional.Bool**| Whether it should be a dessert | 
 **fruit** | **optional.Bool**| Whether it should be a fruit | 
 **herbOrSpice** | **optional.Bool**| Whether it should be an herb or a spice | 
 **sauce** | **optional.Bool**| Whether it should be a sauce | 
 **soup** | **optional.Bool**| Whether it should be a soup | 
 **beverage** | **optional.Bool**| Whether it should be a beverage | 
 **alcoholicBeverage** | **optional.Bool**| Whether it should be an alcoholic beverage | 
 **juice** | **optional.Bool**| Whether it should be a juice | 
 **tea** | **optional.Bool**| Whether it should be a tea | 

### Return type

[**FoodBaseResponse**](FoodBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

