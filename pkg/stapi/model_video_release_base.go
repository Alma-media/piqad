/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Base video release, returned in search results
type VideoReleaseBase struct {
	// Video release unique ID
	Uid string `json:"uid,omitempty"`
	// Video release title
	Title string `json:"title,omitempty"`
	// Series this video release is associated with
	Series *SeriesHeader `json:"series,omitempty"`
	// Season this video release is associated with
	Season *SeasonHeader `json:"season,omitempty"`
	// Video release format
	Format *VideoReleaseFormat `json:"format,omitempty"`
	// Number of episodes
	NumberOfEpisodes int32 `json:"numberOfEpisodes,omitempty"`
	// Number of feature-length episodes
	NumberOfFeatureLengthEpisodes int32 `json:"numberOfFeatureLengthEpisodes,omitempty"`
	// Number of data carriers (like DVD, VCD, VHS etc.)
	NumberOfDataCarriers int32 `json:"numberOfDataCarriers,omitempty"`
	// Run time, in minutes
	RunTime int32 `json:"runTime,omitempty"`
	// Starting year of video release story
	YearFrom int32 `json:"yearFrom,omitempty"`
	// Ending year of video release story
	YearTo int32 `json:"yearTo,omitempty"`
	// Region free release date
	RegionFreeReleaseDate string `json:"regionFreeReleaseDate,omitempty"`
	// Region 1/A release date
	Region1AReleaseDate string `json:"region1AReleaseDate,omitempty"`
	// Region 1 slimline release date
	Region1SlimlineReleaseDate string `json:"region1SlimlineReleaseDate,omitempty"`
	// Region 2/B release date
	Region2BReleaseDate string `json:"region2BReleaseDate,omitempty"`
	// Region 2 slimline release date
	Region2SlimlineReleaseDate string `json:"region2SlimlineReleaseDate,omitempty"`
	// Region 4 release date
	Region4AReleaseDate string `json:"region4AReleaseDate,omitempty"`
	// Region 4 slimline release date
	Region4SlimlineReleaseDate string `json:"region4SlimlineReleaseDate,omitempty"`
	// Whether this video has been release on Amazon.com
	AmazonDigitalRelease bool `json:"amazonDigitalRelease,omitempty"`
	// Whether this video has been release on Dailymotion
	DailymotionDigitalRelease bool `json:"dailymotionDigitalRelease,omitempty"`
	// Whether this video has been release on Google Play
	GooglePlayDigitalRelease bool `json:"googlePlayDigitalRelease,omitempty"`
	// Whether this video has been release on iTunes
	ITunesDigitalRelease bool `json:"iTunesDigitalRelease,omitempty"`
	// Whether this video has been release on UltraViolet
	UltraVioletDigitalRelease bool `json:"ultraVioletDigitalRelease,omitempty"`
	// Whether this video has been release on Vimeo
	VimeoDigitalRelease bool `json:"vimeoDigitalRelease,omitempty"`
	// Whether this video has been release on VUDU
	VuduDigitalRelease bool `json:"vuduDigitalRelease,omitempty"`
	// Whether this video has been release on Xbox SmartGlass
	XboxSmartGlassDigitalRelease bool `json:"xboxSmartGlassDigitalRelease,omitempty"`
	// Whether this video has been release on YouTube
	YouTubeDigitalRelease bool `json:"youTubeDigitalRelease,omitempty"`
	// Whether this video has been release on Netflix
	NetflixDigitalRelease bool `json:"netflixDigitalRelease,omitempty"`
}