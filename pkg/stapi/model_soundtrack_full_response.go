/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Response object for single soundtrack query
type SoundtrackFullResponse struct {
	// Soundtrack, if found
	Soundtrack *SoundtrackFull `json:"soundtrack,omitempty"`
}