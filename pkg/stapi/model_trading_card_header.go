/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Header trading card, embedded in other objects
type TradingCardHeader struct {
	// Trading card unique ID
	Uid string `json:"uid,omitempty"`
	// Trading card name
	Name string `json:"name,omitempty"`
}
