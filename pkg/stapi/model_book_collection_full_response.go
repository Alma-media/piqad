/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for single book collection query
type BookCollectionFullResponse struct {
	// Book collection, if found
	BookCollection *BookCollectionFull `json:"bookCollection,omitempty"`
}
