/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base book series, returned in search results
type BookSeriesBase struct {
	// Book series unique ID
	Uid string `json:"uid,omitempty"`
	// Book series title
	Title string `json:"title,omitempty"`
	// Year from which the book series was published
	PublishedYearFrom int32 `json:"publishedYearFrom,omitempty"`
	// Month from which the book series was published
	PublishedMonthFrom int32 `json:"publishedMonthFrom,omitempty"`
	// Year to which the book series was published
	PublishedYearTo int32 `json:"publishedYearTo,omitempty"`
	// Month to which the book series was published
	PublishedMonthTo int32 `json:"publishedMonthTo,omitempty"`
	// Number of pages
	NumberOfBooks int32 `json:"numberOfBooks,omitempty"`
	// Starting year of book series stories
	YearFrom int32 `json:"yearFrom,omitempty"`
	// Ending year of book series stories
	YearTo int32 `json:"yearTo,omitempty"`
	// Whether it's a miniseries
	Miniseries bool `json:"miniseries,omitempty"`
	// Whether it's a e-book series
	EBookSeries bool `json:"eBookSeries,omitempty"`
}