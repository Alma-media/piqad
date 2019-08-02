/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base episode, returned in search results
type EpisodeBase struct {
	// Episode unique ID
	Uid string `json:"uid,omitempty"`
	// Episode title
	Title string `json:"title,omitempty"`
	// Episode title in German
	TitleGerman string `json:"titleGerman,omitempty"`
	// Episode title in Italian
	TitleItalian string `json:"titleItalian,omitempty"`
	// Episode title in Japanese
	TitleJapanese string `json:"titleJapanese,omitempty"`
	// Series this episode belongs to
	Series *SeriesHeader `json:"series,omitempty"`
	// Season this episode belongs to
	Season *SeasonHeader `json:"season,omitempty"`
	// Season number
	SeasonNumber int32 `json:"seasonNumber,omitempty"`
	// Episode number in season
	EpisodeNumber int32 `json:"episodeNumber,omitempty"`
	// Production serial number
	ProductionSerialNumber string `json:"productionSerialNumber,omitempty"`
	// Whether it's a feature length episode
	FeatureLength bool `json:"featureLength,omitempty"`
	// Starting stardate of episode story
	StardateFrom float32 `json:"stardateFrom,omitempty"`
	// Ending stardate of episode story
	StardateTo float32 `json:"stardateTo,omitempty"`
	// Starting year of episode story
	YearFrom int32 `json:"yearFrom,omitempty"`
	// Ending year of episode story
	YearTo int32 `json:"yearTo,omitempty"`
	// Date the episode was first aired in the United States
	UsAirDate string `json:"usAirDate,omitempty"`
	// Date the episode script was completed
	FinalScriptDate string `json:"finalScriptDate,omitempty"`
}