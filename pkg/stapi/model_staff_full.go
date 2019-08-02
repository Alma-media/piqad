/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full staff, returned when queried using UID
type StaffFull struct {
	// Staff unique ID
	Uid string `json:"uid,omitempty"`
	// Staff name
	Name string `json:"name,omitempty"`
	// Staff birth name
	BirthName string `json:"birthName,omitempty"`
	// Staff gender
	Gender *Gender `json:"gender,omitempty"`
	// Date the staff was born
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	// Place the staff was born
	PlaceOfBirth string `json:"placeOfBirth,omitempty"`
	// Date the staff died
	DateOfDeath string `json:"dateOfDeath,omitempty"`
	// Place the staff died
	PlaceOfDeath string `json:"placeOfDeath,omitempty"`
	// Whether this person is from art department
	ArtDepartment bool `json:"artDepartment,omitempty"`
	// Whether this person is an art director
	ArtDirector bool `json:"artDirector,omitempty"`
	// Whether this person is a production designer
	ProductionDesigner bool `json:"productionDesigner,omitempty"`
	// Whether this person is from camera and electrical department
	CameraAndElectricalDepartment bool `json:"cameraAndElectricalDepartment,omitempty"`
	// Whether this person is a cinematographer
	Cinematographer bool `json:"cinematographer,omitempty"`
	// Whether this person is from casting department
	CastingDepartment bool `json:"castingDepartment,omitempty"`
	// Whether this person is from costume department
	CostumeDepartment bool `json:"costumeDepartment,omitempty"`
	// Whether this person is a custume designer
	CostumeDesigner bool `json:"costumeDesigner,omitempty"`
	// Whether this person is a director
	Director bool `json:"director,omitempty"`
	// Whether this person is an assistant or secound unit director director
	AssistantOrSecondUnitDirector bool `json:"assistantOrSecondUnitDirector,omitempty"`
	// Whether this person is an exhibit and attraction staff
	ExhibitAndAttractionStaff bool `json:"exhibitAndAttractionStaff,omitempty"`
	// Whether this person is a film editor
	FilmEditor bool `json:"filmEditor,omitempty"`
	// Whether this person is a linguist
	Linguist bool `json:"linguist,omitempty"`
	// Whether this person is a location staff
	LocationStaff bool `json:"locationStaff,omitempty"`
	// Whether this person is a make-up staff
	MakeupStaff bool `json:"makeupStaff,omitempty"`
	// Whether this person is from music department
	MusicDepartment bool `json:"musicDepartment,omitempty"`
	// Whether this person is a composer
	Composer bool `json:"composer,omitempty"`
	// Whether this person is a personal assistant
	PersonalAssistant bool `json:"personalAssistant,omitempty"`
	// Whether this person is a producer
	Producer bool `json:"producer,omitempty"`
	// Whether this person is a production associate
	ProductionAssociate bool `json:"productionAssociate,omitempty"`
	// Whether this person is a production staff
	ProductionStaff bool `json:"productionStaff,omitempty"`
	// Whether this person is a publication staff
	PublicationStaff bool `json:"publicationStaff,omitempty"`
	// Whether this person is a science consultant
	ScienceConsultant bool `json:"scienceConsultant,omitempty"`
	// Whether this person is from sound department
	SoundDepartment bool `json:"soundDepartment,omitempty"`
	// Whether this person is a special and visual effects staff
	SpecialAndVisualEffectsStaff bool `json:"specialAndVisualEffectsStaff,omitempty"`
	// Whether this person is an author
	Author bool `json:"author,omitempty"`
	// Whether this person is an audio author
	AudioAuthor bool `json:"audioAuthor,omitempty"`
	// Whether this person is a calendar artist
	CalendarArtist bool `json:"calendarArtist,omitempty"`
	// Whether this person is a comic artist
	ComicArtist bool `json:"comicArtist,omitempty"`
	// Whether this person is a comic author
	ComicAuthor bool `json:"comicAuthor,omitempty"`
	// Whether this person is a comic color artist
	ComicColorArtist bool `json:"comicColorArtist,omitempty"`
	// Whether this person is a comic interior artist
	ComicInteriorArtist bool `json:"comicInteriorArtist,omitempty"`
	// Whether this person is a comic ink artist
	ComicInkArtist bool `json:"comicInkArtist,omitempty"`
	// Whether this person is a comic pencil artist
	ComicPencilArtist bool `json:"comicPencilArtist,omitempty"`
	// Whether this person is a comic letter artist
	ComicLetterArtist bool `json:"comicLetterArtist,omitempty"`
	// Whether this person is a comic strip artist
	ComicStripArtist bool `json:"comicStripArtist,omitempty"`
	// Whether this person is a game artist
	GameArtist bool `json:"gameArtist,omitempty"`
	// Whether this person is a game author
	GameAuthor bool `json:"gameAuthor,omitempty"`
	// Whether this person is a novel artist
	NovelArtist bool `json:"novelArtist,omitempty"`
	// Whether this person is a novel author
	NovelAuthor bool `json:"novelAuthor,omitempty"`
	// Whether this person is a reference artist
	ReferenceArtist bool `json:"referenceArtist,omitempty"`
	// Whether this person is a reference author
	ReferenceAuthor bool `json:"referenceAuthor,omitempty"`
	// Whether this person is a publication artist
	PublicationArtist bool `json:"publicationArtist,omitempty"`
	// Whether this person is a publication designer
	PublicationDesigner bool `json:"publicationDesigner,omitempty"`
	// Whether this person is a publication editor
	PublicationEditor bool `json:"publicationEditor,omitempty"`
	// Whether this person is a publicity artist
	PublicityArtist bool `json:"publicityArtist,omitempty"`
	// Whether this person is a part of CBS digital staff
	CbsDigitalStaff bool `json:"cbsDigitalStaff,omitempty"`
	// Whether this person is a part of ILM production staff
	IlmProductionStaff bool `json:"ilmProductionStaff,omitempty"`
	// Whether this person is a special features artist
	SpecialFeaturesStaff bool `json:"specialFeaturesStaff,omitempty"`
	// Whether this person is a story editor
	StoryEditor bool `json:"storyEditor,omitempty"`
	// Whether this person is a studio executive
	StudioExecutive bool `json:"studioExecutive,omitempty"`
	// Whether this person is from stunt department
	StuntDepartment bool `json:"stuntDepartment,omitempty"`
	// Whether this person is from transportation department
	TransportationDepartment bool `json:"transportationDepartment,omitempty"`
	// Whether this person is video game production staff
	VideoGameProductionStaff bool `json:"videoGameProductionStaff,omitempty"`
	// Whether this person is a writer
	Writer bool `json:"writer,omitempty"`
	// Episodes written by this person
	WrittenEpisodes []EpisodeBase `json:"writtenEpisodes,omitempty"`
	// Episodes to which this person has written teleplay
	TeleplayAuthoredEpisodes []EpisodeBase `json:"teleplayAuthoredEpisodes,omitempty"`
	// Episodes to which this person has written story
	StoryAuthoredEpisodes []EpisodeBase `json:"storyAuthoredEpisodes,omitempty"`
	// Episodes directed by this person
	DirectedEpisodes []EpisodeBase `json:"directedEpisodes,omitempty"`
	// Episodes on which this person worked
	Episodes []EpisodeBase `json:"episodes,omitempty"`
	// Movies written by this person
	WrittenMovies []MovieBase `json:"writtenMovies,omitempty"`
	// Movies to which this person has written screenplay
	ScreenplayAuthoredMovies []MovieBase `json:"screenplayAuthoredMovies,omitempty"`
	// Movies to which this person has written story
	StoryAuthoredMovies []MovieBase `json:"storyAuthoredMovies,omitempty"`
	// Movies directed by this person
	DirectedMovies []MovieBase `json:"directedMovies,omitempty"`
	// Movies produced by this person
	ProducedMovies []MovieBase `json:"producedMovies,omitempty"`
	// Movies on which this person worked
	Movies []MovieBase `json:"movies,omitempty"`
}
