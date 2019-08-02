/*
 * STAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stapi

// Full trading card, returned when queried using UID
type TradingCardFull struct {
	// Trading card unique ID
	Uid string `json:"uid,omitempty"`
	// Trading card name
	Name string `json:"name,omitempty"`
	// Trading card set this card belongs to
	TradingCardSet *TradingCardSetBase `json:"tradingCardSet,omitempty"`
	// Trading card deck this card belongs to
	TradingCardDeck *TradingCardDeckBase `json:"tradingCardDeck,omitempty"`
	// Trading card number
	Number string `json:"number,omitempty"`
	// Release year, if set was releases over multiple years
	ReleaseYear int32 `json:"releaseYear,omitempty"`
	// Production run, if different from trading card set production run
	ProductionRun int32 `json:"productionRun,omitempty"`
}