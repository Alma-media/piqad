/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header spacecraft, embedded in other objects
type SpacecraftHeader struct {
	// Spacecraft unique ID
	Uid string `json:"uid,omitempty"`
	// Spacecraft name
	Name string `json:"name,omitempty"`
}
