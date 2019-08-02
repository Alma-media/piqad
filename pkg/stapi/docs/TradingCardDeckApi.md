# \TradingCardDeckApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TradingCardDeckGet**](TradingCardDeckApi.md#TradingCardDeckGet) | **Get** /tradingCardDeck | 
[**TradingCardDeckSearchGet**](TradingCardDeckApi.md#TradingCardDeckSearchGet) | **Get** /tradingCardDeck/search | 
[**TradingCardDeckSearchPost**](TradingCardDeckApi.md#TradingCardDeckSearchPost) | **Post** /tradingCardDeck/search | 


# **TradingCardDeckGet**
> TradingCardDeckFullResponse TradingCardDeckGet(ctx, uid, optional)


Retrival of a single trading card deck

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| trading card deck unique ID | 
 **optional** | ***TradingCardDeckGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardDeckGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**TradingCardDeckFullResponse**](TradingCardDeckFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradingCardDeckSearchGet**
> TradingCardDeckBaseResponse TradingCardDeckSearchGet(ctx, optional)


Pagination over trading card decks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingCardDeckSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardDeckSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**TradingCardDeckBaseResponse**](TradingCardDeckBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradingCardDeckSearchPost**
> TradingCardDeckBaseResponse TradingCardDeckSearchPost(ctx, optional)


Searching trading card decks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingCardDeckSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingCardDeckSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Trading card deck name | 
 **tradingCardSetUid** | **optional.String**| UID of trading card set | 

### Return type

[**TradingCardDeckBaseResponse**](TradingCardDeckBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

