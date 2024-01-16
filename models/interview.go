package models

type Interview struct {
	LevelOne       bool `json:"levelOne"`
	LevelTwo       bool `json:"levelTwo"`
	Managerial     bool `json:"managerial"`
	OfferRolledOut bool `json:"offerRolledOut"`
	OfferStatus OfferStatus `json:"offerStatus"`
}
