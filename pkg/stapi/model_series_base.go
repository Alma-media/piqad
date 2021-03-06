/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base series, returned in search results
type SeriesBase struct {
	// Series unique ID
	Uid string `json:"uid,omitempty"`
	// Series title
	Title string `json:"title,omitempty"`
	// Series abbreviation
	Abbreviation string `json:"abbreviation,omitempty"`
	// Year the series production started
	ProductionStartYear int32 `json:"productionStartYear,omitempty"`
	// Year the series production ended
	ProductionEndYear int32 `json:"productionEndYear,omitempty"`
	// Date the series originally ran from
	OriginalRunStartDate string `json:"originalRunStartDate,omitempty"`
	// Date the series originally ran to
	OriginalRunEndDate string `json:"originalRunEndDate,omitempty"`
	// Number of seasons
	SeasonsCount int32 `json:"seasonsCount,omitempty"`
	// Number of episodes
	EpisodesCount int32 `json:"episodesCount,omitempty"`
	// Number of feature length episodes
	FeatureLengthEpisodesCount int32 `json:"featureLengthEpisodesCount,omitempty"`
	// Company that produced the series
	ProductionCompany *CompanyHeader `json:"productionCompany,omitempty"`
	// Company that originally broadcasted the series
	OriginalBroadcaster *CompanyHeader `json:"originalBroadcaster,omitempty"`
}
