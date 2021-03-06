/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full species, returned when queried using UID
type SpeciesFull struct {
	// Species unique ID
	Uid string `json:"uid,omitempty"`
	// Species name
	Name string `json:"name,omitempty"`
	// Homeworld of the species
	Homeworld *AstronomicalObjectBase `json:"homeworld,omitempty"`
	// Quadrant the species originates from
	Quadrant *AstronomicalObjectBase `json:"quadrant,omitempty"`
	// Whether it's an extinct species
	ExtinctSpecies bool `json:"extinctSpecies,omitempty"`
	// Whether it's a warp-capable species
	WarpCapableSpecies bool `json:"warpCapableSpecies,omitempty"`
	// Whether it's an extra-galactic species
	ExtraGalacticSpecies bool `json:"extraGalacticSpecies,omitempty"`
	// Whether it's a humanoid species
	HumanoidSpecies bool `json:"humanoidSpecies,omitempty"`
	// Whether it's a reptilian species
	ReptilianSpecies bool `json:"reptilianSpecies,omitempty"`
	// Whether it's a non-corporeal species
	NonCorporealSpecies bool `json:"nonCorporealSpecies,omitempty"`
	// Whether it's a shapeshifting species
	ShapeshiftingSpecies bool `json:"shapeshiftingSpecies,omitempty"`
	// Whether it's a spaceborne species
	SpaceborneSpecies bool `json:"spaceborneSpecies,omitempty"`
	// Whether it's a telepathic species
	TelepathicSpecies bool `json:"telepathicSpecies,omitempty"`
	// Whether it's a trans-dimensional species
	TransDimensionalSpecies bool `json:"transDimensionalSpecies,omitempty"`
	// Whether it's a unnamed species
	UnnamedSpecies bool `json:"unnamedSpecies,omitempty"`
	// Whether this species is from alternate reality
	AlternateReality bool `json:"alternateReality,omitempty"`
	// Characters belonging to the species
	Characters []CharacterBase `json:"characters,omitempty"`
}
