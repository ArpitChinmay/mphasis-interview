package handlers

import (
	"net/http"
	"strconv"

	"github.com/ArpitChinmay/mphasis-interview/models"
	"github.com/ArpitChinmay/mphasis-interview/main"
	"github.com/gin-gonic/gin"
)

func SearchInterviewDetailsByLayerHandler(c *gin.Context) {
	layer := c.Query("layer")
	layerAsInt, err := strconv.Atoi(layer)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "layer provided is invalid"})
		return
	}

	DetailsOfSelectedCandidates := make([]models.Interview, 0)
	DetailsOfRejectedCandidates := make([]models.Interview, 0)

	if layerAsInt == 1 {
		for i := 0; i < len(InterviewDetails); i++ {
			if main.InterviewDetails[i].LevelOne {
				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
			} else {
				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
			}
		}
	}

	if layerAsInt == 2 {
		for i := 0; i < len(InterviewDetails); i++ {
			if InterviewDetails[i].LevelTwo {
				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
			} else {
				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
			}
		}
	}

	if layerAsInt == 3 {
		for i := 0; i < len(InterviewDetails); i++ {
			if InterviewDetails[i].Managerial {
				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
			} else {
				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
			}
		}
	}

	if layerAsInt == 4 {
		for i := 0; i < len(InterviewDetails); i++ {
			if InterviewDetails[i].LevelTwo {
				DetailsOfSelectedCandidates = append(DetailsOfSelectedCandidates, InterviewDetails[i])
			} else {
				DetailsOfRejectedCandidates = append(DetailsOfRejectedCandidates, InterviewDetails[i])
			}
		}
	}
	c.JSON(http.StatusOK, DetailsOfSelectedCandidates)
}
