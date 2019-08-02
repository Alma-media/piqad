# \CompanyApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompanyGet**](CompanyApi.md#CompanyGet) | **Get** /company | 
[**CompanySearchGet**](CompanyApi.md#CompanySearchGet) | **Get** /company/search | 
[**CompanySearchPost**](CompanyApi.md#CompanySearchPost) | **Post** /company/search | 


# **CompanyGet**
> CompanyFullResponse CompanyGet(ctx, uid, optional)


Retrival of a single company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Company unique ID | 
 **optional** | ***CompanyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanyGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**CompanyFullResponse**](CompanyFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompanySearchGet**
> CompanyBaseResponse CompanySearchGet(ctx, optional)


Pagination over companies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanySearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanySearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**CompanyBaseResponse**](CompanyBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompanySearchPost**
> CompanyBaseResponse CompanySearchPost(ctx, optional)


Searching companies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompanySearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompanySearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Company name | 
 **broadcaster** | **optional.Bool**| Whether it should be a broadcaster | 
 **collectibleCompany** | **optional.Bool**| Whether it should be a collectible company | 
 **conglomerate** | **optional.Bool**| Whether it should be a conglomerate | 
 **digitalVisualEffectsCompany** | **optional.Bool**| Whether it should be a digital visual effects company | 
 **distributor** | **optional.Bool**| Whether it should be a distributor | 
 **gameCompany** | **optional.Bool**| Whether it should be a game company | 
 **filmEquipmentCompany** | **optional.Bool**| Whether it should be a film equipment company | 
 **makeUpEffectsStudio** | **optional.Bool**| Whether it should be a make-up effects studio | 
 **mattePaintingCompany** | **optional.Bool**| Whether it should be a matte painting company | 
 **modelAndMiniatureEffectsCompany** | **optional.Bool**| Whether it should be a model and miniature effects company | 
 **postProductionCompany** | **optional.Bool**| Whether it should be a post-production company | 
 **productionCompany** | **optional.Bool**| Whether it should be a production company | 
 **propCompany** | **optional.Bool**| Whether it should be a prop company | 
 **recordLabel** | **optional.Bool**| Whether it should be a record label | 
 **specialEffectsCompany** | **optional.Bool**| Whether it should be a special effects company | 
 **tvAndFilmProductionCompany** | **optional.Bool**| Whether it should be a TV and film production company | 
 **videoGameCompany** | **optional.Bool**| Whether it should be a video game company | 

### Return type

[**CompanyBaseResponse**](CompanyBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

