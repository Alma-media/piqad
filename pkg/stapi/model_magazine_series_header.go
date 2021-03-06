/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header book series, embedded in other objects
type MagazineSeriesHeader struct {
	// Magazine series unique ID
	Uid string `json:"uid,omitempty"`
	// Magazine series title
	Title string `json:"title,omitempty"`
}
