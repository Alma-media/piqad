/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Error object
type ModelError struct {
	// Error code
	Code string `json:"code,omitempty"`
	// Error message
	Message string `json:"message,omitempty"`
}
