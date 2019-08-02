# \PerformerApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformerGet**](PerformerApi.md#PerformerGet) | **Get** /performer | 
[**PerformerSearchGet**](PerformerApi.md#PerformerSearchGet) | **Get** /performer/search | 
[**PerformerSearchPost**](PerformerApi.md#PerformerSearchPost) | **Post** /performer/search | 


# **PerformerGet**
> PerformerFullResponse PerformerGet(ctx, uid, optional)


Retrival of a single performer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Performer unique ID | 
 **optional** | ***PerformerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PerformerGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**PerformerFullResponse**](PerformerFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformerSearchGet**
> PerformerBaseResponse PerformerSearchGet(ctx, optional)


Pagination over performers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PerformerSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PerformerSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**PerformerBaseResponse**](PerformerBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformerSearchPost**
> PerformerBaseResponse PerformerSearchPost(ctx, optional)


Searching performers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PerformerSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PerformerSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Performer name | 
 **birthName** | **optional.String**| Performer birth name | 
 **gender** | **optional.String**| Performer gender | 
 **dateOfBirthFrom** | **optional.String**| Minimal date the performer was born | 
 **dateOfBirthTo** | **optional.String**| Maximal date the performer was born | 
 **placeOfBirth** | **optional.String**| Place the performer was born | 
 **dateOfDeathFrom** | **optional.String**| Minimal date the performer died | 
 **dateOfDeathTo** | **optional.String**| Maximal date the performer died | 
 **placeOfDeath** | **optional.String**| Place the performer died | 
 **animalPerformer** | **optional.Bool**| Whether it should be an animal performer | 
 **disPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: Discovery | 
 **ds9Performer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: Deep Space Nine | 
 **entPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: Enterprise | 
 **filmPerformer** | **optional.Bool**| Whether it should be a performer that appeared in a Star Trek movie | 
 **standInPerformer** | **optional.Bool**| Whether it should be a stand-in performer | 
 **stuntPerformer** | **optional.Bool**| Whether it should be a stunt performer | 
 **tasPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: The Animated Series | 
 **tngPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: The Next Generation | 
 **tosPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: The Original Series | 
 **videoGamePerformer** | **optional.Bool**| Whether it should be a video game performer | 
 **voicePerformer** | **optional.Bool**| Whether it should be a voice performer | 
 **voyPerformer** | **optional.Bool**| Whether it should be a performer that appeared in Star Trek: Voyager | 

### Return type

[**PerformerBaseResponse**](PerformerBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

