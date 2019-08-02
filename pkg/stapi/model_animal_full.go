/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full animal, returned when queried using UID
type AnimalFull struct {
	// Animal unique ID
	Uid string `json:"uid,omitempty"`
	// Animal name
	Name string `json:"name,omitempty"`
	// Whether it's an earth animal
	EarthAnimal bool `json:"earthAnimal,omitempty"`
	// Whether it's an earth insect
	EarthInsect bool `json:"earthInsect,omitempty"`
	// Whether it's an avian
	Avian bool `json:"avian,omitempty"`
	// Whether it's a canine
	Canine bool `json:"canine,omitempty"`
	// Whether it's a feline
	Feline bool `json:"feline,omitempty"`
}
