/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base conflict, returned in search results
type ConflictFull struct {
	// Conflict unique ID
	Uid string `json:"uid,omitempty"`
	// Conflict name
	Name string `json:"name,omitempty"`
	// Starting year of the conflict
	YearFrom int32 `json:"yearFrom,omitempty"`
	// Ending year of the conflict
	YearTo int32 `json:"yearTo,omitempty"`
	// Whether it is an Earth conflict
	EarthConflict bool `json:"earthConflict,omitempty"`
	// Whether this conflict is a part of war involving Federation
	FederationWar bool `json:"federationWar,omitempty"`
	// Whether this conflict is a part of war involving the Klingons
	KlingonWar bool `json:"klingonWar,omitempty"`
	// Whether this conflict is a Dominion war battle
	DominionWarBattle bool `json:"dominionWarBattle,omitempty"`
	// Whether this conflict is from alternate reality
	AlternateReality bool `json:"alternateReality,omitempty"`
	// Locations this conflict occurred at
	Locations []LocationBase `json:"locations,omitempty"`
	// Organization involved in conflict on first side
	FirstSideBelligerents []OrganizationBase `json:"firstSideBelligerents,omitempty"`
	// Commanders involved in conflict on first side
	FirstSideCommanders []CharacterBase `json:"firstSideCommanders,omitempty"`
	// Organization involved in conflict on second side
	SecondSideBelligerents []OrganizationBase `json:"secondSideBelligerents,omitempty"`
	// Commanders involved in conflict on second side
	SecondSideCommanders []CharacterBase `json:"secondSideCommanders,omitempty"`
}
