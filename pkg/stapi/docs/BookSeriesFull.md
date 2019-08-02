# BookSeriesFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Book series unique ID | [optional] [default to null]
**Title** | **string** | Book series title | [optional] [default to null]
**PublishedYearFrom** | **int32** | Year from which the book series was published | [optional] [default to null]
**PublishedMonthFrom** | **int32** | Month from which the book series was published | [optional] [default to null]
**PublishedYearTo** | **int32** | Year to which the book series was published | [optional] [default to null]
**PublishedMonthTo** | **int32** | Month to which the book series was published | [optional] [default to null]
**NumberOfBooks** | **int32** | Number of books in book series | [optional] [default to null]
**YearFrom** | **int32** | Starting year of book series stories | [optional] [default to null]
**YearTo** | **int32** | Ending year of book series stories | [optional] [default to null]
**Miniseries** | **bool** | Whether it&#39;s a miniseries | [optional] [default to null]
**EBookSeries** | **bool** | Whether it&#39;s a e-book series | [optional] [default to null]
**ParentSeries** | [**[]BookSeriesBase**](BookSeriesBase.md) | Book series this book series is included in | [optional] [default to null]
**ChildSeries** | [**[]BookSeriesBase**](BookSeriesBase.md) | Child book series included in this book series | [optional] [default to null]
**Publishers** | [**[]CompanyBase**](CompanyBase.md) | Companies that published this book series | [optional] [default to null]
**Books** | [**[]BookBase**](BookBase.md) | Books included in this book series | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


