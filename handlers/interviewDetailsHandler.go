package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/mphasis-interview/db/datareader"
	"github.com/ArpitChinmay/mphasis-interview/models"
	"github.com/gin-gonic/gin"
)

type InterviewHandler struct {
	context          context.Context
	interviewDetails []models.Interview
	reader           *datareader.DataReader
}

func NewInterviewHandler(c *gin.Context, db *sql.DB) *InterviewHandler {
	return &InterviewHandler{context: c, interviewDetails: make([]models.Interview, 0), reader: new(datareader.DataReader)}
}

func (handler *InterviewHandler) GetAllCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataFromDatabase(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetAllCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataFromDatabaseForLevelTwo(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetAllCandidatesAtLevelThree(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataFromDatabaseForLevelThree(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadSelectedCandidateDataFromDatabaseForLevelOne(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadRejectedCandidateDataFromDatabaseForLevelOne(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadSelectedCandidateDataFromDatabaseForLevelTwo(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadRejectedCandidateDataFromDatabaseForLevelTwo(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtLevelThree(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadSelectedCandidateDataFromDatabaseForLevelThree(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtLevelThree(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadRejectedCandidateDataFromDatabaseForLevelThree(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetOfferedCandidatesThatHaveAccepted(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadOfferedCandidateDataWhoAcceptedOffer(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetOfferedCandidatesThatHaveAcceptanceAwaited(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadOfferedCandidateDataWithAcceptanceAwaited(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetOfferedCandidatesThatHaveAcceptedAndOnboarded(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadOfferedCandidateDataWithOfferAcceptedAndOnboarded(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

