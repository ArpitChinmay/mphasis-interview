package models

type OfferStatus struct {
	// status can be
	// 0. Not applicable
	// 1. awaited
	// 2. acceptance requested
	// 3. accepted
	// 4. onboarded
	OfferId     int    `json:"offerId"`
	Description string `json:"description"`
}
