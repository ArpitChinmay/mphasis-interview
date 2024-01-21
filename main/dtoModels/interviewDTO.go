package dtomodels

import (
	"time"

	"github.com/ArpitChinmay/mphasis-interview/models"
)

type InterviewDTO struct {
	CandidateName  string    `json:"Candidate Name"`
	NameOnResume   string    `json:"Name On Resume"`
	DOB            time.Time `json:"Date Of Birth"`
	PhoneNumber    string    `json:"Phone"`
	City           string    `json:"City"`
	Country        string    `json:"Country"`
	LinkedIn       string    `json:"LinkedIn"`
	Github         string    `json:"Github"`
	Stackoverflow  string    `json:"Stack Overflow"`
	Interests      string    `json:"Interests"`
	LevelOne       string    `json:"First Round"`
	LevelTwo       string    `json:"Second Round"`
	Managerial     string    `json:"Manager Round"`
	OfferRolledOut string    `json:"Offer Rolled Out"`
	OfferStatus    string    `json:"Offer Status"`
}

func (r InterviewDTO) MapInterviewDetails(databaseModel *models.Interview) InterviewDTO {
	result := InterviewDTO{}
	result.CandidateName = databaseModel.CandidateName
	result.NameOnResume = databaseModel.NameOnResume
	result.PhoneNumber = databaseModel.PhoneNumber
	result.City = databaseModel.City
	result.Country = databaseModel.Country
	result.LinkedIn = databaseModel.LinkedIn
	result.Github = databaseModel.Github
	result.Stackoverflow = databaseModel.Stackoverflow
	result.Interests = databaseModel.CandidateInterest
	result.LevelOne = databaseModel.LevelOneStatus
	result.LevelTwo = databaseModel.LevelTwoStatus
	result.Managerial = databaseModel.ManaerialInterviewStatus
	result.OfferStatus = databaseModel.OfferStatus
	
	date, err := time.Parse("YYYY-MM-DD HH:MM:SS", databaseModel.Dob)
	if err != nil {
		result.DOB = time.Time{}
	} else {
		result.DOB = date
	}
	
	if databaseModel.OfferRolledOut == 1 {
		result.OfferRolledOut = "Yes"
	} else {
		result.OfferRolledOut = "No"
	}

	return result
}
