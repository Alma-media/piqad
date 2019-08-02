/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for elements search
type ElementBaseResponse struct {
	// Response page
	Page *ResponsePage `json:"page,omitempty"`
	// Response sort
	Sort *ResponseSort `json:"sort,omitempty"`
	// List of elements matching given criteria
	Elements []ElementBase `json:"elements,omitempty"`
}