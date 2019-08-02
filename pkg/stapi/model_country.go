/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Real-world country
type Country struct {
	// Country unique ID
	Uid string `json:"uid,omitempty"`
	// Country name
	Name string `json:"name,omitempty"`
	// ISO 3166-1 alpha-2 code
	Iso31661Alpha2Code string `json:"iso31661Alpha2Code,omitempty"`
}
