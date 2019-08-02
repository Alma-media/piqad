# \ConflictApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConflictGet**](ConflictApi.md#ConflictGet) | **Get** /conflict | 
[**ConflictSearchGet**](ConflictApi.md#ConflictSearchGet) | **Get** /conflict/search | 
[**ConflictSearchPost**](ConflictApi.md#ConflictSearchPost) | **Post** /conflict/search | 


# **ConflictGet**
> ConflictFullResponse ConflictGet(ctx, uid, optional)


Retrival of a single conflict

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Conflict unique ID | 
 **optional** | ***ConflictGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConflictGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**ConflictFullResponse**](ConflictFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConflictSearchGet**
> ConflictBaseResponse ConflictSearchGet(ctx, optional)


Pagination over conflicts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConflictSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConflictSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**ConflictBaseResponse**](ConflictBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConflictSearchPost**
> ConflictBaseResponse ConflictSearchPost(ctx, optional)


Searching conflicts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConflictSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConflictSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Conflict name | 
 **yearFrom** | **optional.Int32**| Starting year of the conflict | 
 **yearTo** | **optional.Int32**| Ending year of the conflict | 
 **earthConflict** | **optional.Bool**| Whether it should be an Earth conflict | 
 **federationWar** | **optional.Bool**| Whether this conflict should be a part of war involving Federation | 
 **klingonWar** | **optional.Bool**| Whether this conflict should be a part of war involving the Klingons | 
 **dominionWarBattle** | **optional.Bool**| Whether this conflict should be a Dominion war battle | 
 **alternateReality** | **optional.Bool**| Whether this conflict should be from alternate reality | 

### Return type

[**ConflictBaseResponse**](ConflictBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

