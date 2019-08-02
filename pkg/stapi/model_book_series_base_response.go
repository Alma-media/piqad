/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for book series search
type BookSeriesBaseResponse struct {
	// Response page
	Page *ResponsePage `json:"page,omitempty"`
	// Response sort
	Sort *ResponseSort `json:"sort,omitempty"`
	// List of book series matching given criteria
	BookSeries []BookSeriesBase `json:"bookSeries,omitempty"`
}