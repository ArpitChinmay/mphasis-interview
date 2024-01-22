package datareader

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/mphasis-interview/models"
)

const (
	gET_ALL_CANDIDATE_AT_LEVEL_ONE = `SELECT I.interviewId, 
    I.CandidateId, 
    COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
    R.resumeId, 
    R.name AS 'nameOnResume', 
    R.dob, 
    R.workexmapperId, 
    PD.phone, 
    PD.currentcity, 
    PD.country, 
    PD.socialId, 
    S.linkedIn, 
    S.github, 
    S.stackoverflow, 
    PD.skillId, 
    SK.description AS 'candidateSkills', 
    PD.InterestId, 
    INT.description AS 'candidateInterests', 
    PD.educationDetailsId, 
    I.LevelOne, 
    LS1.description AS 'levelOneStatus', 
    I.LevelTwo, 
    LS2.description AS 'LevelTwoStatus', 
    I.Managerial, 
    LS3.description AS 'ManagerialInterviewStatus', 
    I.OfferRolledOut, 
    I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
    INNER JOIN User U
        ON I.CandidateId = U.userId
        AND U.candidateFlag = 1
    INNER JOIN Resume R
        ON I.ResumeId = R.resumeId
    INNER JOIN personalDetails PD
        ON R.personalDetailsId = PD.personaldetailsId
    INNER JOIN socials S
        ON S.socialId = PD.socialId
    INNER JOIN skills SK
        ON SK.skillId = PD.skillId
    INNER JOIN Interests INT
        ON INT.interestId = PD.interestId
    INNER JOIN LayerStatus LS1
        ON I.LevelOne = LS1.layerStatusId
    INNER JOIN LayerStatus LS2
        ON I.LevelTwo = LS2.layerStatusId
    INNER JOIN LayerStatus LS3
        ON I.Managerial = LS3.layerStatusId
    INNER JOIN offerStatus OS
        ON I.OfferStatus = OS.offerId
`

	gET_SELECTED_CANDIDATE_AT_LEVEL_ONE = `SELECT I.interviewId, 
I.CandidateId, 
COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
R.resumeId, 
R.name AS 'nameOnResume', 
R.dob, 
R.workexmapperId, 
PD.phone, 
PD.currentcity, 
PD.country, 
PD.socialId, 
S.linkedIn, 
S.github, 
S.stackoverflow, 
PD.skillId, 
SK.description AS 'candidateSkills', 
PD.InterestId, 
INT.description AS 'candidateInterests', 
PD.educationDetailsId, 
I.LevelOne, 
LS1.description AS 'levelOneStatus', 
I.LevelTwo, 
LS2.description AS 'LevelTwoStatus', 
I.Managerial, 
LS3.description AS 'ManagerialInterviewStatus', 
I.OfferRolledOut, 
I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
INNER JOIN User U
	ON I.CandidateId = U.userId
	AND I.LevelOne = 2
	AND U.candidateFlag = 1
INNER JOIN Resume R
	ON I.ResumeId = R.resumeId
INNER JOIN personalDetails PD
	ON R.personalDetailsId = PD.personaldetailsId
INNER JOIN socials S
	ON S.socialId = PD.socialId
INNER JOIN skills SK
	ON SK.skillId = PD.skillId
INNER JOIN Interests INT
	ON INT.interestId = PD.interestId
INNER JOIN LayerStatus LS1
	ON I.LevelOne = LS1.layerStatusId
INNER JOIN LayerStatus LS2
	ON I.LevelTwo = LS2.layerStatusId
INNER JOIN LayerStatus LS3
	ON I.Managerial = LS3.layerStatusId
INNER JOIN offerStatus OS
	ON I.OfferStatus = OS.offerId
`

	gET_SELECTED_CANDIDATE_AT_LEVEL_TWO = `SELECT I.interviewId, 
I.CandidateId, 
COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
R.resumeId, 
R.name AS 'nameOnResume', 
R.dob, 
R.workexmapperId, 
PD.phone, 
PD.currentcity, 
PD.country, 
PD.socialId, 
S.linkedIn, 
S.github, 
S.stackoverflow, 
PD.skillId, 
SK.description AS 'candidateSkills', 
PD.InterestId, 
INT.description AS 'candidateInterests', 
PD.educationDetailsId, 
I.LevelOne, 
LS1.description AS 'levelOneStatus', 
I.LevelTwo, 
LS2.description AS 'LevelTwoStatus', 
I.Managerial, 
LS3.description AS 'ManagerialInterviewStatus', 
I.OfferRolledOut, 
I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
INNER JOIN User U
	ON I.CandidateId = U.userId
	AND I.LevelTwo = 2
	AND U.candidateFlag = 1
INNER JOIN Resume R
	ON I.ResumeId = R.resumeId
INNER JOIN personalDetails PD
	ON R.personalDetailsId = PD.personaldetailsId
INNER JOIN socials S
	ON S.socialId = PD.socialId
INNER JOIN skills SK
	ON SK.skillId = PD.skillId
INNER JOIN Interests INT
	ON INT.interestId = PD.interestId
INNER JOIN LayerStatus LS1
	ON I.LevelOne = LS1.layerStatusId
INNER JOIN LayerStatus LS2
	ON I.LevelTwo = LS2.layerStatusId
INNER JOIN LayerStatus LS3
	ON I.Managerial = LS3.layerStatusId
INNER JOIN offerStatus OS
	ON I.OfferStatus = OS.offerId
`
	gET_SELECTED_CANDIDATE_AT_LEVEL_THREE = `SELECT I.interviewId, 
I.CandidateId, 
COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
R.resumeId, 
R.name AS 'nameOnResume', 
R.dob, 
R.workexmapperId, 
PD.phone, 
PD.currentcity, 
PD.country, 
PD.socialId, 
S.linkedIn, 
S.github, 
S.stackoverflow, 
PD.skillId, 
SK.description AS 'candidateSkills', 
PD.InterestId, 
INT.description AS 'candidateInterests', 
PD.educationDetailsId, 
I.LevelOne, 
LS1.description AS 'levelOneStatus', 
I.LevelTwo, 
LS2.description AS 'LevelTwoStatus', 
I.Managerial, 
LS3.description AS 'ManagerialInterviewStatus', 
I.OfferRolledOut, 
I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
INNER JOIN User U
	ON I.CandidateId = U.userId
	AND I.Managerial= 2
	AND U.candidateFlag = 1
INNER JOIN Resume R
	ON I.ResumeId = R.resumeId
INNER JOIN personalDetails PD
	ON R.personalDetailsId = PD.personaldetailsId
INNER JOIN socials S
	ON S.socialId = PD.socialId
INNER JOIN skills SK
	ON SK.skillId = PD.skillId
INNER JOIN Interests INT
	ON INT.interestId = PD.interestId
INNER JOIN LayerStatus LS1
	ON I.LevelOne = LS1.layerStatusId
INNER JOIN LayerStatus LS2
	ON I.LevelTwo = LS2.layerStatusId
INNER JOIN LayerStatus LS3
	ON I.Managerial = LS3.layerStatusId
INNER JOIN offerStatus OS
	ON I.OfferStatus = OS.offerId`

	gET_REJECTED_CANDIDATE_AT_LEVEL_ONE = `SELECT I.interviewId, 
    I.CandidateId, 
    COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
    R.resumeId, 
    R.name AS 'nameOnResume', 
    R.dob, 
    R.workexmapperId, 
    PD.phone, 
    PD.currentcity, 
    PD.country, 
    PD.socialId, 
    S.linkedIn, 
    S.github, 
    S.stackoverflow, 
    PD.skillId, 
    SK.description AS 'candidateSkills', 
    PD.InterestId, 
    INT.description AS 'candidateInterests', 
    PD.educationDetailsId, 
    I.LevelOne, 
    LS1.description AS 'levelOneStatus', 
    I.LevelTwo, 
    LS2.description AS 'LevelTwoStatus', 
    I.Managerial, 
    LS3.description AS 'ManagerialInterviewStatus', 
    I.OfferRolledOut, 
    I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
    INNER JOIN User U
        ON I.CandidateId = U.userId
		AND I.LevelOne = 3
        AND U.candidateFlag = 1
    INNER JOIN Resume R
        ON I.ResumeId = R.resumeId
    INNER JOIN personalDetails PD
        ON R.personalDetailsId = PD.personaldetailsId
    INNER JOIN socials S
        ON S.socialId = PD.socialId
    INNER JOIN skills SK
        ON SK.skillId = PD.skillId
    INNER JOIN Interests INT
        ON INT.interestId = PD.interestId
    INNER JOIN LayerStatus LS1
        ON I.LevelOne = LS1.layerStatusId
    INNER JOIN LayerStatus LS2
        ON I.LevelTwo = LS2.layerStatusId
    INNER JOIN LayerStatus LS3
        ON I.Managerial = LS3.layerStatusId
    INNER JOIN offerStatus OS
        ON I.OfferStatus = OS.offerId`

	gET_REJECTED_CANDIDATE_AT_LEVEL_TWO = `SELECT I.interviewId, 
		I.CandidateId, 
		COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
		R.resumeId, 
		R.name AS 'nameOnResume', 
		R.dob, 
		R.workexmapperId, 
		PD.phone, 
		PD.currentcity, 
		PD.country, 
		PD.socialId, 
		S.linkedIn, 
		S.github, 
		S.stackoverflow, 
		PD.skillId, 
		SK.description AS 'candidateSkills', 
		PD.InterestId, 
		INT.description AS 'candidateInterests', 
		PD.educationDetailsId, 
		I.LevelOne, 
		LS1.description AS 'levelOneStatus', 
		I.LevelTwo, 
		LS2.description AS 'LevelTwoStatus', 
		I.Managerial, 
		LS3.description AS 'ManagerialInterviewStatus', 
		I.OfferRolledOut, 
		I.OfferStatus, OS.description AS 'offerStatus'
	FROM Interview I
		INNER JOIN User U
			ON I.CandidateId = U.userId
			AND I.LevelTwo = 3
			AND U.candidateFlag = 1
		INNER JOIN Resume R
			ON I.ResumeId = R.resumeId
		INNER JOIN personalDetails PD
			ON R.personalDetailsId = PD.personaldetailsId
		INNER JOIN socials S
			ON S.socialId = PD.socialId
		INNER JOIN skills SK
			ON SK.skillId = PD.skillId
		INNER JOIN Interests INT
			ON INT.interestId = PD.interestId
		INNER JOIN LayerStatus LS1
			ON I.LevelOne = LS1.layerStatusId
		INNER JOIN LayerStatus LS2
			ON I.LevelTwo = LS2.layerStatusId
		INNER JOIN LayerStatus LS3
			ON I.Managerial = LS3.layerStatusId
		INNER JOIN offerStatus OS
			ON I.OfferStatus = OS.offerId`

	gET_REJECTED_CANDIDATE_AT_LEVEL_THREE = `SELECT I.interviewId, 
	I.CandidateId, 
	COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
	R.resumeId, 
	R.name AS 'nameOnResume', 
	R.dob, 
	R.workexmapperId, 
	PD.phone, 
	PD.currentcity, 
	PD.country, 
	PD.socialId, 
	S.linkedIn, 
	S.github, 
	S.stackoverflow, 
	PD.skillId, 
	SK.description AS 'candidateSkills', 
	PD.InterestId, 
	INT.description AS 'candidateInterests', 
	PD.educationDetailsId, 
	I.LevelOne, 
	LS1.description AS 'levelOneStatus', 
	I.LevelTwo, 
	LS2.description AS 'LevelTwoStatus', 
	I.Managerial, 
	LS3.description AS 'ManagerialInterviewStatus', 
	I.OfferRolledOut, 
	I.OfferStatus, OS.description AS 'offerStatus'
FROM Interview I
	INNER JOIN User U
		ON I.CandidateId = U.userId
		AND I.Managerial = 3
		AND U.candidateFlag = 1
	INNER JOIN Resume R
		ON I.ResumeId = R.resumeId
	INNER JOIN personalDetails PD
		ON R.personalDetailsId = PD.personaldetailsId
	INNER JOIN socials S
		ON S.socialId = PD.socialId
	INNER JOIN skills SK
		ON SK.skillId = PD.skillId
	INNER JOIN Interests INT
		ON INT.interestId = PD.interestId
	INNER JOIN LayerStatus LS1
		ON I.LevelOne = LS1.layerStatusId
	INNER JOIN LayerStatus LS2
		ON I.LevelTwo = LS2.layerStatusId
	INNER JOIN LayerStatus LS3
		ON I.Managerial = LS3.layerStatusId
	INNER JOIN offerStatus OS
		ON I.OfferStatus = OS.offerId`

	gET_CANDIDATE_OFFERED_POSITION_ACCEPTED = `SELECT I.interviewId, 
		I.CandidateId, 
		COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
		R.resumeId, 
		R.name AS 'nameOnResume', 
		R.dob, 
		R.workexmapperId, 
		PD.phone, 
		PD.currentcity, 
		PD.country, 
		PD.socialId, 
		S.linkedIn, 
		S.github, 
		S.stackoverflow, 
		PD.skillId, 
		SK.description AS 'candidateSkills', 
		PD.InterestId, 
		INT.description AS 'candidateInterests', 
		PD.educationDetailsId, 
		I.LevelOne, 
		LS1.description AS 'levelOneStatus', 
		I.LevelTwo, 
		LS2.description AS 'LevelTwoStatus', 
		I.Managerial, 
		LS3.description AS 'ManagerialInterviewStatus', 
		I.OfferRolledOut, 
		I.OfferStatus, OS.description AS 'offerStatus'
	FROM Interview I
		INNER JOIN User U
			ON I.CandidateId = U.userId
			AND U.candidateFlag = 1
			AND I.OfferRolledOut = 1
			AND I.OfferStatus = 2
		INNER JOIN Resume R
			ON I.ResumeId = R.resumeId
		INNER JOIN personalDetails PD
			ON R.personalDetailsId = PD.personaldetailsId
		INNER JOIN socials S
			ON S.socialId = PD.socialId
		INNER JOIN skills SK
			ON SK.skillId = PD.skillId
		INNER JOIN Interests INT
			ON INT.interestId = PD.interestId
		INNER JOIN LayerStatus LS1
			ON I.LevelOne = LS1.layerStatusId
		INNER JOIN LayerStatus LS2
			ON I.LevelTwo = LS2.layerStatusId
		INNER JOIN LayerStatus LS3
			ON I.Managerial = LS3.layerStatusId
		INNER JOIN offerStatus OS
			ON I.OfferStatus = OS.offerId`

	gET_CANDIDATE_OFFERED_POSITION_ACCEPTENCE_AWAITED = `SELECT I.interviewId, 
	I.CandidateId, 
	COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
	R.resumeId, 
	R.name AS 'nameOnResume', 
	R.dob, 
	R.workexmapperId, 
	PD.phone, 
	PD.currentcity, 
	PD.country, 
	PD.socialId, 
	S.linkedIn, 
	S.github, 
	S.stackoverflow, 
	PD.skillId, 
	SK.description AS 'candidateSkills', 
	PD.InterestId, 
	INT.description AS 'candidateInterests', 
	PD.educationDetailsId, 
	I.LevelOne, 
	LS1.description AS 'levelOneStatus', 
	I.LevelTwo, 
	LS2.description AS 'LevelTwoStatus', 
	I.Managerial, 
	LS3.description AS 'ManagerialInterviewStatus', 
	I.OfferRolledOut, 
	I.OfferStatus, OS.description AS 'offerStatus'
	FROM Interview I
		INNER JOIN User U
			ON I.CandidateId = U.userId
			AND U.candidateFlag = 1
			AND I.OfferRolledOut = 1
			AND I.OfferStatus = 1
		INNER JOIN Resume R
			ON I.ResumeId = R.resumeId
		INNER JOIN personalDetails PD
			ON R.personalDetailsId = PD.personaldetailsId
		INNER JOIN socials S
			ON S.socialId = PD.socialId
		INNER JOIN skills SK
			ON SK.skillId = PD.skillId
		INNER JOIN Interests INT
			ON INT.interestId = PD.interestId
		INNER JOIN LayerStatus LS1
			ON I.LevelOne = LS1.layerStatusId
		INNER JOIN LayerStatus LS2
			ON I.LevelTwo = LS2.layerStatusId
		INNER JOIN LayerStatus LS3
			ON I.Managerial = LS3.layerStatusId
		INNER JOIN offerStatus OS
			ON I.OfferStatus = OS.offerId`

	gET_CANDIDATE_OFFERED_POSITION_ACCEPTED_ONBOARDED = `SELECT I.interviewId, 
	I.CandidateId, 
	COALESCE(U.title, '') || U.firstName || ' ' || COALESCE(U.midname, '') || ' ' || U.lastName as 'candidateName', 
	R.resumeId, 
	R.name AS 'nameOnResume', 
	R.dob, 
	R.workexmapperId, 
	PD.phone, 
	PD.currentcity, 
	PD.country, 
	PD.socialId, 
	S.linkedIn, 
	S.github, 
	S.stackoverflow, 
	PD.skillId, 
	SK.description AS 'candidateSkills', 
	PD.InterestId, 
	INT.description AS 'candidateInterests', 
	PD.educationDetailsId, 
	I.LevelOne, 
	LS1.description AS 'levelOneStatus', 
	I.LevelTwo, 
	LS2.description AS 'LevelTwoStatus', 
	I.Managerial, 
	LS3.description AS 'ManagerialInterviewStatus', 
	I.OfferRolledOut, 
	I.OfferStatus, OS.description AS 'offerStatus'
	FROM Interview I
		INNER JOIN User U
			ON I.CandidateId = U.userId
			AND U.candidateFlag = 1
			AND I.OfferRolledOut = 1
			AND I.OfferStatus = 3
		INNER JOIN Resume R
			ON I.ResumeId = R.resumeId
		INNER JOIN personalDetails PD
			ON R.personalDetailsId = PD.personaldetailsId
		INNER JOIN socials S
			ON S.socialId = PD.socialId
		INNER JOIN skills SK
			ON SK.skillId = PD.skillId
		INNER JOIN Interests INT
			ON INT.interestId = PD.interestId
		INNER JOIN LayerStatus LS1
			ON I.LevelOne = LS1.layerStatusId
		INNER JOIN LayerStatus LS2
			ON I.LevelTwo = LS2.layerStatusId
		INNER JOIN LayerStatus LS3
			ON I.Managerial = LS3.layerStatusId
		INNER JOIN offerStatus OS
			ON I.OfferStatus = OS.offerId`
)

type DataReader struct {
	database         *sql.DB
	interviewDetails []models.Interview
}

func NewDataReader(db *sql.DB) *DataReader {
	return &DataReader{database: db, interviewDetails: make([]models.Interview, 0)}
}

func (datareader *DataReader) ReadInterviewDataFromDatabase(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_ALL_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataFromDatabaseForLevelTwo(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataFromDatabaseForLevelThree(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadSelectedCandidateDataFromDatabaseForLevelOne(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadSelectedCandidateDataFromDatabaseForLevelTwo(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadSelectedCandidateDataFromDatabaseForLevelThree(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_THREE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadRejectedCandidateDataFromDatabaseForLevelOne(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadRejectedCandidateDataFromDatabaseForLevelTwo(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadRejectedCandidateDataFromDatabaseForLevelThree(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_THREE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadOfferedCandidateDataWhoAcceptedOffer(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_CANDIDATE_OFFERED_POSITION_ACCEPTED)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadOfferedCandidateDataWithAcceptanceAwaited(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_CANDIDATE_OFFERED_POSITION_ACCEPTENCE_AWAITED)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadOfferedCandidateDataWithOfferAcceptedAndOnboarded(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_CANDIDATE_OFFERED_POSITION_ACCEPTED_ONBOARDED)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewId, &candidate.CandidateId, &candidate.CandidateName,
			&candidate.ResumeId, &candidate.NameOnResume, &candidate.Dob, &candidate.WorkExMapperId,
			&candidate.PhoneNumber, &candidate.City, &candidate.Country, &candidate.SocialId, &candidate.LinkedIn,
			&candidate.Github, &candidate.Stackoverflow, &candidate.SkillId, &candidate.CandidateSkills, &candidate.InterestId,
			&candidate.CandidateInterest, &candidate.EducationDetailsId, &candidate.LevelOneId, &candidate.LevelOneStatus,
			&candidate.LevelTwoId, &candidate.LevelTwoStatus, &candidate.ManaerialInterviewId, &candidate.ManaerialInterviewStatus,
			&candidate.OfferRolledOut, &candidate.OfferStatusId, &candidate.OfferStatus)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}
