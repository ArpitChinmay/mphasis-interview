package models

import "time"

type Resume struct {
	ResumeId        int             `json:"resumeId"`
	CandidateId     int             `json:"candidateId"`
	Name            string          `json:"candidateName"`
	Dob             time.Time       `json:"dateofbirth"`
	WordEx          Workex          `json:"WorkEx"`
	PersonalDetails PersonalDetails `json:"personalDetails"`
	ResumeStatus    bool            `json:"resumeStatus"` // Resume status can be accepted or rejected
}
