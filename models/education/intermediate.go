package models

type Intermediate struct {
	CandidateId int     `json:"candidateId"`
	School_name string  `json:"schoolName"`
	Board       string  `json:"board"`
	Percentage  float64 `json:"percentage"`
	Year        int     `json:"passingYear"`
}
