package models

type Interview struct {
	CandidateId    string `json:"candidateId"`
	ResumeId       string `json:"resumeId"`
	LevelOne       int    `json:"levelOne"`
	LevelTwo       int    `json:"levelTwo"`
	Managerial     int    `json:"managerial"`
	OfferRolledOut int    `json:"offerRolledOut"`
	OfferStatus    int    `json:"offerStatus"`
}
