/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for single comic series query
type ComicSeriesFullResponse struct {
	// Comic series, if found
	ComicSeries *ComicSeriesFull `json:"comicSeries,omitempty"`
}
