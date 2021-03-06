/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header astronomical object, embedded in other objects
type AstronomicalObjectHeader struct {
	// Astronomical object's unique ID
	Uid string `json:"uid,omitempty"`
	// Astronomical object name
	Name string `json:"name,omitempty"`
}
