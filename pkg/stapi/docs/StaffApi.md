# \StaffApi

All URIs are relative to *http://stapi.co/api/v1/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StaffGet**](StaffApi.md#StaffGet) | **Get** /staff | 
[**StaffSearchGet**](StaffApi.md#StaffSearchGet) | **Get** /staff/search | 
[**StaffSearchPost**](StaffApi.md#StaffSearchPost) | **Post** /staff/search | 


# **StaffGet**
> StaffFullResponse StaffGet(ctx, uid, optional)


Retrival of a single staff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Staff unique ID | 
 **optional** | ***StaffGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StaffGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| API key | 

### Return type

[**StaffFullResponse**](StaffFullResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StaffSearchGet**
> StaffBaseResponse StaffSearchGet(ctx, optional)


Pagination over staff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StaffSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StaffSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **apiKey** | **optional.String**| API key | 

### Return type

[**StaffBaseResponse**](StaffBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StaffSearchPost**
> StaffBaseResponse StaffSearchPost(ctx, optional)


Searching staff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StaffSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StaffSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Zero-based page number | 
 **pageSize** | **optional.Int32**| Page size | 
 **sort** | **optional.String**| Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC | 
 **apiKey** | **optional.String**| API key | 
 **name** | **optional.String**| Staff name | 
 **birthName** | **optional.String**| Staff birth name | 
 **gender** | **optional.String**| Staff gender | 
 **dateOfBirthFrom** | **optional.String**| Minimal date the staff was born | 
 **dateOfBirthTo** | **optional.String**| Maximal date the staff was born | 
 **placeOfBirth** | **optional.String**| Place the staff was born | 
 **dateOfDeathFrom** | **optional.String**| Minimal date the staff died | 
 **dateOfDeathTo** | **optional.String**| Maximal date the staff died | 
 **placeOfDeath** | **optional.String**| Place the staff died | 
 **artDepartment** | **optional.Bool**| Whether this person should be from art department | 
 **artDirector** | **optional.Bool**| Whether this person should be an art director | 
 **productionDesigner** | **optional.Bool**| Whether this person should be a production designer | 
 **cameraAndElectricalDepartment** | **optional.Bool**| Whether this person should be from camera and electrical department | 
 **cinematographer** | **optional.Bool**| Whether this person should be a cinematographer | 
 **castingDepartment** | **optional.Bool**| Whether this person should be from casting department | 
 **costumeDepartment** | **optional.Bool**| Whether this person should be from costume department | 
 **costumeDesigner** | **optional.Bool**| Whether this person should be a custume designer | 
 **director** | **optional.Bool**| Whether this person should be a director | 
 **assistantOrSecondUnitDirector** | **optional.Bool**| Whether this person should be an assistant or secound unit director director | 
 **exhibitAndAttractionStaff** | **optional.Bool**| Whether this person should be an exhibit and attraction staff | 
 **filmEditor** | **optional.Bool**| Whether this person should be a film editor | 
 **linguist** | **optional.Bool**| Whether this person should be a linguist | 
 **locationStaff** | **optional.Bool**| Whether this person should be a location staff | 
 **makeupStaff** | **optional.Bool**| Whether this person should be a make-up staff | 
 **musicDepartment** | **optional.Bool**| Whether this person should be from music department | 
 **composer** | **optional.Bool**| Whether this person should be a composer | 
 **personalAssistant** | **optional.Bool**| Whether this person should be a personal assistant | 
 **producer** | **optional.Bool**| Whether this person should be a producer | 
 **productionAssociate** | **optional.Bool**| Whether this person should be a production associate | 
 **productionStaff** | **optional.Bool**| Whether this person should be a production staff | 
 **publicationStaff** | **optional.Bool**| Whether this person should be a publication staff | 
 **scienceConsultant** | **optional.Bool**| Whether this person should be a science consultant | 
 **soundDepartment** | **optional.Bool**| Whether this person should be from sound department | 
 **specialAndVisualEffectsStaff** | **optional.Bool**| Whether this person should be a special and visual effects staff | 
 **author** | **optional.Bool**| Whether this person should be an author | 
 **audioAuthor** | **optional.Bool**| Whether this person should be an audio author | 
 **calendarArtist** | **optional.Bool**| Whether this person should be a calendar artist | 
 **comicArtist** | **optional.Bool**| Whether this person should be a comic artist | 
 **comicAuthor** | **optional.Bool**| Whether this person should be a comic author | 
 **comicColorArtist** | **optional.Bool**| Whether this person should be a comic color artist | 
 **comicInteriorArtist** | **optional.Bool**| Whether this person should be a comic interior artist | 
 **comicInkArtist** | **optional.Bool**| Whether this person should be a comic ink artist | 
 **comicPencilArtist** | **optional.Bool**| Whether this person should be a comic pencil artist | 
 **comicLetterArtist** | **optional.Bool**| Whether this person should be a comic letter artist | 
 **comicStripArtist** | **optional.Bool**| Whether this person should be a comic strip artist | 
 **gameArtist** | **optional.Bool**| Whether this person should be a game artist | 
 **gameAuthor** | **optional.Bool**| Whether this person should be a game author | 
 **novelArtist** | **optional.Bool**| Whether this person should be a novel artist | 
 **novelAuthor** | **optional.Bool**| Whether this person should be a novel author | 
 **referenceArtist** | **optional.Bool**| Whether this person should be a reference artist | 
 **referenceAuthor** | **optional.Bool**| Whether this person should be a reference author | 
 **publicationArtist** | **optional.Bool**| Whether this person should be a publication artist | 
 **publicationDesigner** | **optional.Bool**| Whether this person should be a publication designer | 
 **publicationEditor** | **optional.Bool**| Whether this person should be a publication editor | 
 **publicityArtist** | **optional.Bool**| Whether this person should be a publicity artist | 
 **cbsDigitalStaff** | **optional.Bool**| Whether this person should be a part of CBS digital staff | 
 **ilmProductionStaff** | **optional.Bool**| Whether this person should be a part of ILM production staff | 
 **specialFeaturesStaff** | **optional.Bool**| Whether this person should be a special features artist | 
 **storyEditor** | **optional.Bool**| Whether this person should be a story editor | 
 **studioExecutive** | **optional.Bool**| Whether this person should be a studio executive | 
 **stuntDepartment** | **optional.Bool**| Whether this person should be from stunt department | 
 **transportationDepartment** | **optional.Bool**| Whether this person should be from transportation department | 
 **videoGameProductionStaff** | **optional.Bool**| Whether this person is video game production staff | 
 **writer** | **optional.Bool**| Whether this person should be a writer | 

### Return type

[**StaffBaseResponse**](StaffBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

