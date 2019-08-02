/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for companies search
type CompanyBaseResponse struct {
	// Response page
	Page *ResponsePage `json:"page,omitempty"`
	// Response sort
	Sort *ResponseSort `json:"sort,omitempty"`
	// List of companies matching given criteria
	Companies []CompanyBase `json:"companies,omitempty"`
}
