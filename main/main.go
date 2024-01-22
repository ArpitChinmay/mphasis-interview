package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/ArpitChinmay/mphasis-interview/handlers"
	dtomodels "github.com/ArpitChinmay/mphasis-interview/main/dtoModels"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type OfferStatus struct {
	// status can be
	// 0. Not applicable
	// 1. acceptance awaited
	// 2. accepted
	// 3. onboarded
	OfferId     int    `json:"offerId"`
	Description string `json:"description"`
}

// var OfferStatusDetails []OfferStatus

var DB *sql.DB

var interviewHandler *handlers.InterviewHandler

func main() {
	router := gin.Default()
	router.GET("/interview/:level/", GetCandidatesInterviewDetails)
	router.GET("/interview/:level/search", GetSpecificCandidatesAtLevel)
	router.GET("/interview/offer", GetCandidatesAtOfferLevel)
	router.Run(":5000")
}

// Seed the temporary data for Interview details and offer status on starting the application.
func init() {
	db, err := sql.Open("sqlite3", "../db/candidateData")
	if err != nil {
		log.Println("could not connect to the database...")
		log.Fatal(err)
	}
	log.Println("Connected to database...")
	DB = db
	interviewHandler = new(handlers.InterviewHandler)
	log.Println("created interviewHandler object", interviewHandler)
}

func GetCandidatesInterviewDetails(c *gin.Context) {
	level, err := strconv.ParseInt(c.Param("level"), 0, 32)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading level..."})
	}
	// Show the list of all candidates with interview level 1 selected and rejected.
	if level == 1 {
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelOne(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 2 {
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelTwo(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 3 {
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelThree(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "feature not implemented yet"})
	}
}

func GetSpecificCandidatesAtLevel(c *gin.Context) {
	level, err := strconv.ParseInt(c.Param("level"), 0, 32)
	selected, err2 := strconv.ParseBool(c.Query("selected"))
	count, err3 := strconv.ParseBool(c.Query("count"))
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading url params..."})
	}
	if level == 1 {
		if selected {
			detailsOfCandidatesDTO, datacount, err := getSelectedCandidateInterviewDetailsAtLevelOne(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		} else {
			detailsOfCandidatesDTO, datacount, err := getRejectedCandidateInterviewDetailsAtLevelOne(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		}

	} else if level == 2 {
		if selected {
			detailsOfCandidatesDTO, datacount, err := getSelectedCandidateInterviewDetailsAtLevelTwo(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		} else {
			detailsOfCandidatesDTO, datacount, err := getRejectedCandidateInterviewDetailsAtLevelTwo(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		}
	} else if level == 3 {
		if selected {
			detailsOfCandidatesDTO, datacount, err := getSelectedCandidateInterviewDetailsAtLevelThree(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		} else {
			detailsOfCandidatesDTO, datacount, err := getRejectedCandidateInterviewDetailsAtLevelThree(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			if !count {
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			} else {
				c.JSON(http.StatusOK, datacount)
			}
		}
	}

}

func GetCandidatesAtOfferLevel(c *gin.Context) {
	accepted, err := strconv.ParseInt(c.Query("accepted"), 0, 32);
	count, err2 := strconv.ParseBool(c.Query("count"))
	
	if err != nil || err2 != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading url params..."})
	}
	if accepted == 1 {
			detailsOfCandidatesDTO, datacount, err := getOfferedCandidatesThatHaveAccepted(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		if !count {
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			c.JSON(http.StatusOK, datacount)
		} 
	} else if accepted == 2 {
		// show the list of all candidates with offer status "acceptance awaited"
		detailsOfCandidatesDTO, datacount, err := getOfferedCandidatesThatHaveAcceptanceAwaited(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		if !count {
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			c.JSON(http.StatusOK, datacount)
		} 
	} else if accepted == 3 {
		detailsOfCandidatesDTO, datacount, err := getOfferedCandidatesThatHaveAcceptedAndOnboarded(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		if !count {
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			c.JSON(http.StatusOK, datacount)
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
}

func getCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetAllCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetAllCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetAllCandidatesAtLevelThree(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtLevelThree(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtLevelThree(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getOfferedCandidatesThatHaveAccepted(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetOfferedCandidatesThatHaveAccepted(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getOfferedCandidatesThatHaveAcceptanceAwaited(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetOfferedCandidatesThatHaveAcceptanceAwaited(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getOfferedCandidatesThatHaveAcceptedAndOnboarded(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetOfferedCandidatesThatHaveAcceptedAndOnboarded(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}