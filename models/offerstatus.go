package models

type OfferStatus struct {
	OfferId     int    `json:"OfferId"`
	Description string `json:"description"` // status can be acceptance requested, awaited, accepted, onboarded
}
