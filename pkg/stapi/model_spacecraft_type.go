/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Rating of video release, etc.
type SpacecraftType struct {
	// Spacecraft type unique ID
	Uid string `json:"uid,omitempty"`
	// Spacecraft type name
	Name string `json:"name,omitempty"`
}