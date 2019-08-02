/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header weapon, embedded in other objects
type WeaponHeader struct {
	// Weapon unique ID
	Uid string `json:"uid,omitempty"`
	// Weapon name
	Name string `json:"name,omitempty"`
}
