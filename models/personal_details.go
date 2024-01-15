package models

import "github.com/ArpitChinmay/mphasis-interview/models"

type PersonalDetails struct {
	PersonalDetailsId  int                     `json:"personalDetailsId"`
	CandidateId        int                     `json:"candidateId"`
	PhoneNumber        string                  `json:"phonenumber"`
	CurrentCity        string                  `json:"currentCity"`
	CurrentCountry     string                  `json:"currentCountry"`
	Socials            Socials                 `json:"socials"`
	CandidateSkills    []string                `json:"skills"`
	CandidateInterests []string                `json:"interests"`
	EducationDetails   models.EducationDetails `json:"educationDetails"`
}
