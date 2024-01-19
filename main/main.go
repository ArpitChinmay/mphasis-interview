package main

import (
	"database/sql"
	"log"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Interview struct {
	CandidateId    string `json:"candidateId"`
	ResumeId       string `json:"resumeId"`
	LevelOne       int    `json:"levelOne"`
	LevelTwo       int    `json:"levelTwo"`
	Managerial     int    `json:"managerial"`
	OfferRolledOut int    `json:"offerRolledOut"`
	OfferStatus    int    `json:"offerStatus"`
}

type OfferStatus struct {
	// status can be
	// 0. Not applicable
	// 1. acceptance awaited
	// 2. accepted
	// 3. onboarded
	OfferId     int    `json:"offerId"`
	Description string `json:"description"`
}

var InterviewDetails []Interview = make([]Interview, 0)

// var OfferStatusDetails []OfferStatus

var DB *sql.DB

func main() {
	router := gin.Default()
	router.GET("/interview/level1", GetAllCandidatesLevelOne)
	// router.GET("/interview/level2", GetAllCandidatesLevelTwo)
	// router.GET("/interview/managerial", GetAllCandidatesManagerial)
	// router.GET("/interview/level1/search", GetCandidatesLevelOne)
	// router.GET("/interview/level2/search", GetCandidatesLevelTwo)
	// router.GET("/interview/managerial/search", GetCandidatesManagerial)
	// router.GET("/interview/level2/count/search", GetLevelTwoCount)
	// router.GET("/interview/managerial/count/search", GetManagerialCount)
	// router.GET("/interview/offer/", GetCandidatesOfferedPostion)
	router.Run(":5000")
}

// Seed the temporary data for Interview details and offer status on starting the application.
func init() {
	// InterviewDetails = make([]Interview, 0)
	// OfferStatusDetails = make([]OfferStatus, 0)
	// file, _ := os.ReadFile("interview-details")
	// _ = json.Unmarshal([]byte(file), &InterviewDetails)
	// file, _ = os.ReadFile("offerstatus")
	// _ = json.Unmarshal([]byte(file), &OfferStatusDetails)
	db, err := sql.Open("sqlite3", "../db/candidateData")
	if err != nil {
		log.Println("could not connect to the database...")
		log.Fatal(err)
	}
	log.Println("Connected to database...")
	DB = db
	rows, err := DB.Query("SELECT CandidateId, ResumeId, LevelOne, LevelTwo, Managerial, OfferRolledOut, OfferStatus FROM interview")

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
	}

	for rows.Next() {
		candidate := Interview{}
		err = rows.Scan(&candidate.CandidateId, &candidate.ResumeId, &candidate.LevelOne, &candidate.LevelTwo, &candidate.Managerial, &candidate.OfferRolledOut, &candidate.OfferStatus)
		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
		}
		InterviewDetails = append(InterviewDetails, candidate)
	}
	defer rows.Close()
}

// Show the list of all candidates with interview level 1 selected and rejected.
func GetAllCandidatesLevelOne(c *gin.Context) {
	DetailsOfAllCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
	}
	c.JSON(http.StatusOK, DetailsOfAllCandidates)
}

// // show the list of all candidates at level 2 selected or rejected
// func GetAllCandidatesLevelTwo(c *gin.Context) {
// 	DetailsOfAllCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne {
// 			DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
// 		}
// 	}
// 	c.JSON(http.StatusOK, DetailsOfAllCandidates)
// }

// // show the list of all candidates at manager level, selected or rejected
// func GetAllCandidatesManagerial(c *gin.Context) {
// 	DetailsOfAllCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && InterviewDetails[i].Managerial {
// 			DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
// 		}
// 	}
// 	c.JSON(http.StatusOK, DetailsOfAllCandidates)
// }

// // Show the list of candidates either selected or rejected in level one.
// func GetCandidatesLevelOne(c *gin.Context) {
// 	selected, err := strconv.ParseBool(c.Query("selected"))
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
// 		return
// 	}

// 	if selected {
// 		DetailsOfSelectedCandidates := GetDetailsOfSelectedCandidatesLevelOne()
// 		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
// 	} else {
// 		DetailsOfRejectedCandidates := GetDetailsOfRejectedCandidatesLevelOne()
// 		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
// 	}
// }

// // Show the list of candidates either selected or rejected in level two.
// func GetCandidatesLevelTwo(c *gin.Context) {
// 	selected, err := strconv.ParseBool(c.Query("selected"))
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
// 	}

// 	if selected {
// 		DetailsOfSelectedCandidates, _ := GetDetailsOfSelectedCandidatesLevelTwo()
// 		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
// 	} else {
// 		DetailsOfRejectedCandidates, _ := GetDetailsOfRejectedCandidatesLevelTwo()
// 		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
// 	}
// }

// // show the list of candidates either selected or rejected in managerial round
// func GetCandidatesManagerial(c *gin.Context) {
// 	selected, err := strconv.ParseBool(c.Query("selected"))
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
// 		return
// 	}

// 	if selected {
// 		DetailsOfSelectedCandidates, _ := GetDetailsOfSelectedCandidatesManagerialRound()
// 		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
// 	} else {
// 		DetailsOfRejectedCandidates, _ := GetDetailsOfRejectedCandidatesManagerialRound()
// 		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
// 	}
// }

// // Show the count of candidates selected or rejected in level two
// func GetLevelTwoCount(c *gin.Context) {
// 	selected, err := strconv.ParseBool(c.Query("selected"))
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
// 	}

// 	if selected {
// 		_, count := GetDetailsOfSelectedCandidatesLevelTwo()
// 		c.JSON(http.StatusOK, count)
// 	} else {
// 		_, count := GetDetailsOfRejectedCandidatesLevelTwo()
// 		c.JSON(http.StatusOK, count)
// 	}
// }

// // show the count of candidates selected or rejected in managerial round
// func GetManagerialCount(c *gin.Context) {
// 	selected, err := strconv.ParseBool(c.Query("selected"))
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
// 	}

// 	if selected {
// 		_, count := GetDetailsOfSelectedCandidatesManagerialRound()
// 		c.JSON(http.StatusOK, count)
// 	} else {
// 		_, count := GetDetailsOfRejectedCandidatesManagerialRound()
// 		c.JSON(http.StatusOK, count)
// 	}
// }

// // show the candidates who are offered position.
// func GetCandidatesOfferedPostion(c *gin.Context) {
// 	selected, err_acpt := strconv.ParseInt(c.Query("accepted"), 0, 32)
// 	count, err_count := strconv.ParseBool(c.Query("count"))
// 	if err_acpt != nil || err_count != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading acceptance feild or count feild"})
// 	}

// 	// User is requesting for a count of candidates
// 	if count {
// 		// offer acceptance awaited
// 		if selected == 1 {
// 			_, count := GetOfferedCandidatesDetailsWithAwaitedAcceptance()
// 			c.JSON(http.StatusOK, count)
// 		}

// 		// offer accepted.
// 		if selected == 2 {
// 			_, count := GetOfferedCandidatedDetailsWithOfferAccepted()
// 			c.JSON(http.StatusOK, count)
// 		}

// 		// onboarded candidate.
// 		if selected == 3 {
// 			_, count := GetOfferedCandidateDetailsWhoAreOnboarded()
// 			c.JSON(http.StatusOK, count)
// 		}
// 		// User is requesting for details of candidates to whom offer is rolled out.
// 	} else {
// 		// offer acceptance awaited
// 		if selected == 1 {
// 			OfferedCandidates, _ := GetOfferedCandidatesDetailsWithAwaitedAcceptance()
// 			c.JSON(http.StatusOK, OfferedCandidates)
// 		}

// 		// offer accepted.
// 		if selected == 2 {
// 			OfferedCandidates, _ := GetOfferedCandidatedDetailsWithOfferAccepted()
// 			c.JSON(http.StatusOK, OfferedCandidates)
// 		}

// 		// onboarded candidate
// 		if selected == 3 {
// 			OfferedCandidate, _ := GetOfferedCandidateDetailsWhoAreOnboarded()
// 			c.JSON(http.StatusOK, OfferedCandidate)
// 		}
// 	}

// }

// func GetDetailsOfSelectedCandidatesLevelOne() []Interview {
// 	DetailsOfSelectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne {
// 			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfSelectedCandidates
// }

// func GetDetailsOfRejectedCandidatesLevelOne() []Interview {
// 	DetailsOfRejectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if !InterviewDetails[i].LevelOne {
// 			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfRejectedCandidates
// }

// func GetDetailsOfSelectedCandidatesLevelTwo() ([]Interview, int) {
// 	DetailsOfSelectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo {
// 			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfSelectedCandidates, len(DetailsOfSelectedCandidates)
// }

// func GetDetailsOfRejectedCandidatesLevelTwo() ([]Interview, int) {
// 	DetailsOfRejectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne && !InterviewDetails[i].LevelTwo {
// 			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfRejectedCandidates, len(DetailsOfRejectedCandidates)
// }

// func GetDetailsOfSelectedCandidatesManagerialRound() ([]Interview, int) {
// 	DetailsOfSelectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && InterviewDetails[i].Managerial {
// 			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfSelectedCandidates, len(DetailsOfSelectedCandidates)
// }

// func GetDetailsOfRejectedCandidatesManagerialRound() ([]Interview, int) {
// 	DetailsOfRejectedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && !InterviewDetails[i].Managerial {
// 			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfRejectedCandidates, len(DetailsOfRejectedCandidates)
// }

// func GetOfferedCandidatesDetailsWithAwaitedAcceptance() ([]Interview, int) {
// 	DetailsOfOfferedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].OfferRolledOut && InterviewDetails[i].OfferStatus == 1 {
// 			DetailsOfOfferedCandidates = append(DetailsOfOfferedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfOfferedCandidates, len(DetailsOfOfferedCandidates)
// }

// func GetOfferedCandidatedDetailsWithOfferAccepted() ([]Interview, int) {
// 	DetailsOfOfferedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].OfferRolledOut && InterviewDetails[i].OfferStatus == 2 {
// 			DetailsOfOfferedCandidates = append(DetailsOfOfferedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfOfferedCandidates, len(DetailsOfOfferedCandidates)
// }

// func GetOfferedCandidateDetailsWhoAreOnboarded() ([]Interview, int) {
// 	DetailsOfOfferedCandidates := make([]Interview, 0)
// 	for i := 0; i < len(InterviewDetails); i++ {
// 		if InterviewDetails[i].OfferRolledOut && InterviewDetails[i].OfferStatus == 3 {
// 			DetailsOfOfferedCandidates = append(DetailsOfOfferedCandidates, InterviewDetails[i])
// 		}
// 	}
// 	return DetailsOfOfferedCandidates, len(DetailsOfOfferedCandidates)
// }
