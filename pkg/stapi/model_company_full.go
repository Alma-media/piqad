/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full company, returned when queried using UID
type CompanyFull struct {
	// Company unique ID
	Uid string `json:"uid,omitempty"`
	// Company name
	Name string `json:"name,omitempty"`
	// Whether it's a broadcaster
	Broadcaster bool `json:"broadcaster,omitempty"`
	// Whether it's a collectible company
	CollectibleCompany bool `json:"collectibleCompany,omitempty"`
	// Whether it's a conglomerate
	Conglomerate bool `json:"conglomerate,omitempty"`
	// Whether it's a digital visual effects company
	DigitalVisualEffectsCompany bool `json:"digitalVisualEffectsCompany,omitempty"`
	// Whether it's a distributor
	Distributor bool `json:"distributor,omitempty"`
	// Whether it's a game company
	GameCompany bool `json:"gameCompany,omitempty"`
	// Whether it's a film equipment company
	FilmEquipmentCompany bool `json:"filmEquipmentCompany,omitempty"`
	// Whether it's a make-up effects studio
	MakeUpEffectsStudio bool `json:"makeUpEffectsStudio,omitempty"`
	// Whether it's a matte painting company
	MattePaintingCompany bool `json:"mattePaintingCompany,omitempty"`
	// Whether it's a model and miniature effects company
	ModelAndMiniatureEffectsCompany bool `json:"modelAndMiniatureEffectsCompany,omitempty"`
	// Whether it's a post-production company
	PostProductionCompany bool `json:"postProductionCompany,omitempty"`
	// Whether it's a production company
	ProductionCompany bool `json:"productionCompany,omitempty"`
	// Whether it's a prop company
	PropCompany bool `json:"propCompany,omitempty"`
	// Whether it's a record label
	RecordLabel bool `json:"recordLabel,omitempty"`
	// Whether it's a special effects company
	SpecialEffectsCompany bool `json:"specialEffectsCompany,omitempty"`
	// Whether it's a TV and film production company
	TvAndFilmProductionCompany bool `json:"tvAndFilmProductionCompany,omitempty"`
	// Whether it's a video game company
	VideoGameCompany bool `json:"videoGameCompany,omitempty"`
}
