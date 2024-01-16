package main

import (
	// "encoding/json"
	"net/http"
	"strconv"

	// "os"

	"github.com/gin-gonic/gin"
)

type Interview struct {
	CandidateId    int  `json:"candidateId"`
	ResumeId       int  `json:"resumeId"`
	LevelOne       bool `json:"levelOne"`
	LevelTwo       bool `json:"levelTwo"`
	Managerial     bool `json:"managerial"`
	OfferRolledOut bool `json:"offerRolledOut"`
	OfferStatus    int  `json:"offerStatus"`
}

type OfferStatus struct {
	// status can be
	// 0. Not applicable
	// 1. awaited
	// 2. acceptance requested
	// 3. accepted
	// 4. onboarded
	OfferId     int    `json:"offerId"`
	Description string `json:"description"`
}

var InterviewDetails []Interview = []Interview{
	{
		CandidateId:    100000,
		ResumeId:       100000,
		LevelOne:       true,
		LevelTwo:       true,
		Managerial:     true,
		OfferRolledOut: true,
		OfferStatus:    2,
	},

	{
		CandidateId:    100001,
		ResumeId:       100001,
		LevelOne:       true,
		LevelTwo:       true,
		Managerial:     true,
		OfferRolledOut: true,
		OfferStatus:    2,
	},

	{
		CandidateId:    100002,
		ResumeId:       100002,
		LevelOne:       true,
		LevelTwo:       true,
		Managerial:     true,
		OfferRolledOut: false,
		OfferStatus:    1,
	},

	{
		CandidateId:    100003,
		ResumeId:       100003,
		LevelOne:       true,
		LevelTwo:       true,
		Managerial:     false,
		OfferRolledOut: false,
		OfferStatus:    0,
	},
	{
		CandidateId:    100004,
		ResumeId:       100004,
		LevelOne:       true,
		LevelTwo:       false,
		Managerial:     false,
		OfferRolledOut: false,
		OfferStatus:    0,
	},
	{
		CandidateId:    100005,
		ResumeId:       100005,
		LevelOne:       false,
		LevelTwo:       false,
		Managerial:     false,
		OfferRolledOut: false,
		OfferStatus:    0,
	},

	{
		CandidateId:    100006,
		ResumeId:       100006,
		LevelOne:       true,
		LevelTwo:       false,
		Managerial:     false,
		OfferRolledOut: false,
		OfferStatus:    2,
	},

	{
		CandidateId:    100007,
		ResumeId:       100007,
		LevelOne:       false,
		LevelTwo:       false,
		Managerial:     false,
		OfferRolledOut: false,
		OfferStatus:    2,
	},
}
var OfferStatusDetails []OfferStatus

func main() {
	router := gin.Default()
	router.GET("/interview/level1", GetAllCandidatesLevelOne)
	router.GET("/interview/level2", GetAllCandidatesLevelTwo)
	router.GET("/interview/managerial", GetAllCandidatesManagerial)
	router.GET("/interview/level1/search", GetCandidatesLevelOne)
	router.GET("/interview/level2/search", GetCandidatesLevelTwo)
	router.GET("/interview/managerial/search", GetCandidatesManagerial)
	router.GET("/interview/level2/count/search", GetLevelTwoCount)
	router.GET("/interview/managerial/count/search", GetManagerialCount)
	router.Run(":5000")
}

// Seed the temporary data for Interview details and offer status on starting the application.
// func init() {
// 	InterviewDetails = make([]Interview, 0)
// 	OfferStatusDetails = make([]OfferStatus, 0)
// 	file, _ := os.ReadFile("interview-details")
// 	_ = json.Unmarshal([]byte(file), &InterviewDetails)
// 	file, _ = os.ReadFile("offerstatus")
// 	_ = json.Unmarshal([]byte(file), &OfferStatusDetails)
// }

// Show the list of all candidates with interview level 1 selected and rejected.
func GetAllCandidatesLevelOne(c *gin.Context) {
	DetailsOfAllCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
	}
	c.JSON(http.StatusOK, DetailsOfAllCandidates)
}

// show the list of all candidates at level 2 selected or rejected
func GetAllCandidatesLevelTwo(c *gin.Context) {
	DetailsOfAllCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne {
			DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
		}
	}
	c.JSON(http.StatusOK, DetailsOfAllCandidates)
}

// show the list of all candidates at manager level, selected or rejected
func GetAllCandidatesManagerial(c *gin.Context) {
	DetailsOfAllCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && InterviewDetails[i].Managerial {
			DetailsOfAllCandidates = append(DetailsOfAllCandidates, InterviewDetails[i])
		}
	}
	c.JSON(http.StatusOK, DetailsOfAllCandidates)
}

// Show the list of candidates either selected or rejected in level one.
func GetCandidatesLevelOne(c *gin.Context) {
	selected, err := strconv.ParseBool(c.Query("selected"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
		return
	}

	if selected {
		DetailsOfSelectedCandidates := GetDetailsOfSelectedCandidatesLevelOne()
		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
	} else {
		DetailsOfRejectedCandidates := GetDetailsOfRejectedCandidatesLevelOne()
		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
	}
}

// Show the list of candidates either selected or rejected in level two.
func GetCandidatesLevelTwo(c *gin.Context) {
	selected, err := strconv.ParseBool(c.Query("selected"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
	}

	if selected {
		DetailsOfSelectedCandidates, _ := GetDetailsOfSelectedCandidatesLevelTwo()
		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
	} else {
		DetailsOfRejectedCandidates, _ := GetDetailsOfRejectedCandidatesLevelTwo()
		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
	}
}

// show the list of candidates either selected or rejected in managerial round
func GetCandidatesManagerial(c *gin.Context) {
	selected, err := strconv.ParseBool(c.Query("selected"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
		return
	}

	if selected {
		DetailsOfSelectedCandidates, _ := GetDetailsOfSelectedCandidatesManagerialRound()
		c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
	} else {
		DetailsOfRejectedCandidates, _ := GetDetailsOfRejectedCandidatesManagerialRound()
		c.JSON(http.StatusOK, DetailsOfRejectedCandidates)
	}
}

// Show the count of candidates selected or rejected in level two
func GetLevelTwoCount(c *gin.Context) {
	selected, err := strconv.ParseBool(c.Query("selected"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
	}

	if selected {
		_, count := GetDetailsOfSelectedCandidatesLevelTwo()
		c.JSON(http.StatusOK, count)
	} else {
		_, count := GetDetailsOfRejectedCandidatesLevelTwo()
		c.JSON(http.StatusOK, count)
	}
}

// show the count of candidates selected or rejected in managerial round
func GetManagerialCount(c *gin.Context) {
	selected, err := strconv.ParseBool(c.Query("selected"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading choice..."})
	}

	if selected {
		_, count := GetDetailsOfSelectedCandidatesManagerialRound();
		c.JSON(http.StatusOK, count)
	} else {
		_, count := GetDetailsOfRejectedCandidatesManagerialRound();
		c.JSON(http.StatusOK, count)
	}
}

func GetDetailsOfSelectedCandidatesLevelOne() []Interview {
	DetailsOfSelectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne {
			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfSelectedCandidates
}

func GetDetailsOfRejectedCandidatesLevelOne() []Interview {
	DetailsOfRejectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if !InterviewDetails[i].LevelOne {
			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfRejectedCandidates
}

func GetDetailsOfSelectedCandidatesLevelTwo() ([]Interview, int) {
	DetailsOfSelectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo {
			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfSelectedCandidates, len(DetailsOfSelectedCandidates)
}

func GetDetailsOfRejectedCandidatesLevelTwo() ([]Interview, int) {
	DetailsOfRejectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne && !InterviewDetails[i].LevelTwo {
			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfRejectedCandidates, len(DetailsOfRejectedCandidates)
}

func GetDetailsOfSelectedCandidatesManagerialRound() ([]Interview, int) {
	DetailsOfSelectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && InterviewDetails[i].Managerial {
			DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfSelectedCandidates, len(DetailsOfSelectedCandidates)
}

func GetDetailsOfRejectedCandidatesManagerialRound() ([]Interview, int) {
	DetailsOfRejectedCandidates := make([]Interview, 0)
	for i := 0; i < len(InterviewDetails); i++ {
		if InterviewDetails[i].LevelOne && InterviewDetails[i].LevelTwo && !InterviewDetails[i].Managerial {
			DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
		}
	}
	return DetailsOfRejectedCandidates, len(DetailsOfRejectedCandidates)
}

// func SearchInterviewDetailsByLayerHandler(c *gin.Context) {
// 	layer := c.Query("layer")
// 	layerAsInt, err := strconv.Atoi(layer)
// 	if err != nil {
// 		c.JSON(http.StatusNoContent, gin.H{"error": "layer provided is invalid"})
// 		return
// 	}

// 	DetailsOfSelectedCandidates := make([]Interview, 0)
// 	DetailsOfRejectedCandidates := make([]Interview, 0)

// 	if layerAsInt == 1 {
// 		for i := 0; i < len(InterviewDetails); i++ {
// 			if InterviewDetails[i].LevelOne {
// 				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 			} else {
// 				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 			}
// 		}
// 	}

// 	if layerAsInt == 2 {
// 		for i := 0; i < len(InterviewDetails); i++ {
// 			if InterviewDetails[i].LevelTwo {
// 				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 			} else {
// 				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 			}
// 		}
// 	}

// 	if layerAsInt == 3 {
// 		for i := 0; i < len(InterviewDetails); i++ {
// 			if InterviewDetails[i].Managerial {
// 				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 			} else {
// 				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 			}
// 		}
// 	}

// 	if layerAsInt == 4 {
// 		for i := 0; i < len(InterviewDetails); i++ {
// 			if InterviewDetails[i].LevelTwo {
// 				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
// 			} else {
// 				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
// 			}
// 		}
// 	}
// 	c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
// }
