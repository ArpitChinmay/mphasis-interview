package main

import (
	"encoding/json"
	"os"

	mphasisinterview "github.com/ArpitChinmay/mphasis-interview/handlers"
	"github.com/ArpitChinmay/mphasis-interview/models"
	"github.com/gin-gonic/gin"
)

var InterviewDetails []models.Interview
var OfferStatus []models.OfferStatus

func main() {
	router := gin.Default()
	router.GET("/interview", mphasisinterview.SearchInterviewDetailsByLayerHandler)
	router.Run(":5000")
}

// Seed the temporary data for Interview details and offer status on starting the application.
func init() {
	InterviewDetails = make([]models.Interview, 0)
	OfferStatus = make([]models.OfferStatus, 0)
	file, _ := os.ReadFile("interview-details")
	_ = json.Unmarshal([]byte(file), &InterviewDetails)
	file, _ = os.ReadFile("offerstatus")
	_ = json.Unmarshal([]byte(file), &OfferStatus)
}
