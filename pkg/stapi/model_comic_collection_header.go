/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header comic collection, embedded in other objects
type ComicCollectionHeader struct {
	// Comic collection unique ID
	Uid string `json:"uid,omitempty"`
	// Comic collection title
	Title string `json:"title,omitempty"`
}
