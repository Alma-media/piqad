# \TradingCardSetApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TradingCardSetGet**](TradingCardSetApi.md#TradingCardSetGet) | **Get** /tradingCardSet | 
[**TradingCardSetSearchGet**](TradingCardSetApi.md#TradingCardSetSearchGet) | **Get** /tradingCardSet/search | 
[**TradingCardSetSearchPost**](TradingCardSetApi.md#TradingCardSetSearchPost) | **Post** /tradingCardSet/search | 


# **TradingCardSetGet**
> TradingCardSetFullResponse TradingCardSetGet(ctx, uid, optional)


Retrival of a single trading card set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Trading card set unique ID | 
 **optional** | ***TradingCardSetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardSetGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**TradingCardSetFullResponse**](TradingCardSetFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradingCardSetSearchGet**
> TradingCardSetBaseResponse TradingCardSetSearchGet(ctx, optional)


Pagination over trading card sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingCardSetSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardSetSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**TradingCardSetBaseResponse**](TradingCardSetBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradingCardSetSearchPost**
> TradingCardSetBaseResponse TradingCardSetSearchPost(ctx, optional)


Searching trading card sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingCardSetSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardSetSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Trading card set name | 
 **releaseYearFrom** | **optional.Int32**| Starting year the trading card set was released | 
 **releaseYearTo** | **optional.Int32**| Ending year the trading card set was released | 
 **cardsPerPackFrom** | **optional.Int32**| Minimal number of cards per deck | 
 **cardsPerPackTo** | **optional.Int32**| Minimal number of cards per deck | 
 **packsPerBoxFrom** | **optional.Int32**| Minimal number of packs per box | 
 **packsPerBoxTo** | **optional.Int32**| Minimal number of packs per box | 
 **boxesPerCaseFrom** | **optional.Int32**| Minimal number of boxes per case | 
 **boxesPerCaseTo** | **optional.Int32**| Minimal number of boxes per case | 
 **productionRunFrom** | **optional.Int32**| Minimal production run | 
 **productionRunTo** | **optional.Int32**| Minimal production run | 
 **productionRunUnit** | **optional.String**| Production run unit | 
 **cardWidthFrom** | **optional.Float64**| Minimal card width, in inches | 
 **cardWidthTo** | **optional.Float64**| Minimal card width, in inches | 
 **cardHeightFrom** | **optional.Float64**| Minimal card height, in inches | 
 **cardHeightTo** | **optional.Float64**| Minimal card height, in inches | 

### Return type

[**TradingCardSetBaseResponse**](TradingCardSetBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

