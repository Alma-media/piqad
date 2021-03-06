/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base magazine, returned in search results
type MagazineBase struct {
	// Magazine unique ID
	Uid string `json:"uid,omitempty"`
	// Magazine title
	Title string `json:"title,omitempty"`
	// Year the magazine was published
	PublishedYear int32 `json:"publishedYear,omitempty"`
	// Month the magazine was published
	PublishedMonth int32 `json:"publishedMonth,omitempty"`
	// Day the magazine was published
	PublishedDay int32 `json:"publishedDay,omitempty"`
	// Cover publication year
	CoverYear int32 `json:"coverYear,omitempty"`
	// Cover publication month
	CoverMonth int32 `json:"coverMonth,omitempty"`
	// Cover publication day
	CoverDay int32 `json:"coverDay,omitempty"`
	// Number of pages
	NumberOfPages int32 `json:"numberOfPages,omitempty"`
	// Magazine issue number
	IssueNumber string `json:"issueNumber,omitempty"`
}
