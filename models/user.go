package models

type User struct {
	UserId        int       `json:"UserId"`
	CandidateFlag bool      `json:"candidateFlag"`
	Resume        Resume    `json:"Resume"`
	Interview     Interview `json:"InterviewDetails"`
}
