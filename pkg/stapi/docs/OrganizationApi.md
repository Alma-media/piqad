# \OrganizationApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGet**](OrganizationApi.md#OrganizationGet) | **Get** /organization | 
[**OrganizationSearchGet**](OrganizationApi.md#OrganizationSearchGet) | **Get** /organization/search | 
[**OrganizationSearchPost**](OrganizationApi.md#OrganizationSearchPost) | **Post** /organization/search | 


# **OrganizationGet**
> OrganizationFullResponse OrganizationGet(ctx, uid, optional)


Retrival of a single organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Organization unique ID | 
 **optional** | ***OrganizationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**OrganizationFullResponse**](OrganizationFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationSearchGet**
> OrganizationBaseResponse OrganizationSearchGet(ctx, optional)


Pagination over organizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**OrganizationBaseResponse**](OrganizationBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationSearchPost**
> OrganizationBaseResponse OrganizationSearchPost(ctx, optional)


Searching organizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Organization name | 
 **government** | **optional.Bool**| Whether it should be a government | 
 **intergovernmentalOrganization** | **optional.Bool**| Whether it should be an intergovernmental organization | 
 **researchOrganization** | **optional.Bool**| Whether it should be a research organization | 
 **sportOrganization** | **optional.Bool**| Whether it should be a sport organization | 
 **medicalOrganization** | **optional.Bool**| Whether it should be a medical organization | 
 **militaryOrganization** | **optional.Bool**| Whether it should be a military organization | 
 **militaryUnit** | **optional.Bool**| Whether it should be a military unit | 
 **governmentAgency** | **optional.Bool**| Whether it should be a government agency | 
 **lawEnforcementAgency** | **optional.Bool**| Whether it should be a law enforcement agency | 
 **prisonOrPenalColony** | **optional.Bool**| Whether it should be a prison or penal colony | 
 **mirror** | **optional.Bool**| Whether this organization should be from mirror universe | 
 **alternateReality** | **optional.Bool**| Whether this organization should be from alternate reality | 

### Return type

[**OrganizationBaseResponse**](OrganizationBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

