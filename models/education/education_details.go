package models

type EducationDetails struct {
	EducationId  int          `json:"educationId"`
	CandidateId  int          `json:"candidateId"`
	Highschool   Highschool   `json:"highschool_details"`
	Intermediate Intermediate `json:"intermediate_details"`
	Graduation   Graduation   `json:"graduation_details"`
	PostGrad     PostGrad     `json:"postgrad_details"`
	Doctorate    Doctorate    `json:"doctrate_details"`
}
