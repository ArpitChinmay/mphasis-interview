package models

type Interview struct {
	InterviewId              string
	CandidateId              string
	CandidateName            string
	ResumeId                 string
	NameOnResume             string
	Dob                      string
	WorkExMapperId           string
	PhoneNumber              string
	City                     string
	Country                  string
	SocialId                 string
	LinkedIn                 string
	Github                   string
	Stackoverflow            string
	SkillId                  string
	CandidateSkills          string
	InterestId               string
	CandidateInterest        string
	EducationDetailsId       string
	LevelOneId               int
	LevelOneStatus           string
	LevelTwoId               int
	LevelTwoStatus           string
	ManaerialInterviewId     int
	ManaerialInterviewStatus string
	OfferRolledOut           int
	OfferStatusId            int
	OfferStatus              string
}

func (interview *Interview) NewInterview() *Interview {
	return &Interview{}
}
