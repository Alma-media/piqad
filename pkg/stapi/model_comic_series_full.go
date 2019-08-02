/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full comic series, returned when queried using UID
type ComicSeriesFull struct {
	// Comic series unique ID
	Uid string `json:"uid,omitempty"`
	// Comic series title
	Title string `json:"title,omitempty"`
	// Year from which the comic series was published
	PublishedYearFrom int32 `json:"publishedYearFrom,omitempty"`
	// Month from which the comic series was published
	PublishedMonthFrom int32 `json:"publishedMonthFrom,omitempty"`
	// Day from which the comic series was published
	PublishedDayFrom int32 `json:"publishedDayFrom,omitempty"`
	// Year to which the comic series was published
	PublishedYearTo int32 `json:"publishedYearTo,omitempty"`
	// Month to which the comic series was published
	PublishedMonthTo int32 `json:"publishedMonthTo,omitempty"`
	// Day to which the comic series was published
	PublishedDayTo int32 `json:"publishedDayTo,omitempty"`
	// Number of issues
	NumberOfIssues int32 `json:"numberOfIssues,omitempty"`
	// Starting stardate of comic series stories
	StardateFrom float32 `json:"stardateFrom,omitempty"`
	// Ending stardate of comic series stories
	StardateTo float32 `json:"stardateTo,omitempty"`
	// Starting year of comic series stories
	YearFrom int32 `json:"yearFrom,omitempty"`
	// Ending year of comic series stories
	YearTo int32 `json:"yearTo,omitempty"`
	// Whether it's a miniseries
	Miniseries bool `json:"miniseries,omitempty"`
	// Whether it's a photonovel series
	PhotonovelSeries bool `json:"photonovelSeries,omitempty"`
	// Comic series this comic series is included in
	ParentSeries []ComicSeriesBase `json:"parentSeries,omitempty"`
	// Child comic series included in this comic series
	ChildSeries []ComicSeriesBase `json:"childSeries,omitempty"`
	// Companies that published this comic series
	Publishers []CompanyBase `json:"publishers,omitempty"`
	// Comics included in this comic series
	Comics []ComicsBase `json:"comics,omitempty"`
}
