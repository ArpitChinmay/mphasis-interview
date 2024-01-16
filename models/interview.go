package models

type Interview struct {
	CandidateId    int  `json:"candidateId"`
	ResumeId       int  `json:"resumeId"`
	LevelOne       bool `json:"levelOne"`
	LevelTwo       bool `json:"levelTwo"`
	Managerial     bool `json:"managerial"`
	OfferRolledOut bool `json:"offerRolledOut"`
	OfferStatus    int  `json:"offerStatus"`
}
