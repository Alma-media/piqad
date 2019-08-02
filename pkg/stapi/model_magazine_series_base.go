/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base magazine series, returned in search results
type MagazineSeriesBase struct {
	// Magazine series unique ID
	Uid string `json:"uid,omitempty"`
	// Magazine series title
	Title string `json:"title,omitempty"`
	// Year from which the magazine series was published
	PublishedYearFrom int32 `json:"publishedYearFrom,omitempty"`
	// Month from which the magazine series was published
	PublishedMonthFrom int32 `json:"publishedMonthFrom,omitempty"`
	// Year to which the magazine series was published
	PublishedYearTo int32 `json:"publishedYearTo,omitempty"`
	// Month to which the magazine series was published
	PublishedMonthTo int32 `json:"publishedMonthTo,omitempty"`
	// Number of issues
	NumberOfIssues int32 `json:"numberOfIssues,omitempty"`
}
