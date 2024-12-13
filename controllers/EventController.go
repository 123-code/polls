package controllers

import (
	"fmt"
	"math"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EventImpact struct {
	CandidateID   uint      `json:"candidate_id"`
	Province      string    `json:"province"`
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
	Severity      int       `json:"severity"`
	VotesBefore   int64     `json:"votes_before"`
	VotesAfter    int64     `json:"votes_after"`
	ImpactRatio   float64   `json:"impact_ratio"`
}

type EventImpactByProvince struct {
	CandidateID       uint          `json:"candidate_id"`
	Province          string        `json:"province"`
	TotalEvents       int           `json:"total_events"`
	EventDetails      []EventImpact `json:"event_details"`
	TotalVotesBefore  int64         `json:"total_votes_before"`
	TotalVotesAfter   int64         `json:"total_votes_after"`
	VoteImpactRatio   float64       `json:"vote_impact_ratio"`
}

func AnalyzeEventImpactByProvince(c *gin.Context) {
	candidateIDStr := c.Query("candidate_id")
	if candidateIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "candidate_id is required"})
		return
	}

	candidateID, err := strconv.ParseUint(candidateIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate_id"})
		return
	}

	var events []models.Event
	if err := initializers.DB.Where("candidate_id = ?", candidateID).Order("date").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	provinceVotes := make(map[string]int64)
	for _, provinceName := range provinces {
		provinceVotes[provinceName] = 0
	}

	var votes []models.MyVote
	if err := initializers.DB.Where("candidate_id = ?", candidateID).Find(&votes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve votes"})
		return
	}

	for _, vote := range votes {
		fmt.Printf("Vote Province: %s\n", vote.Province)

		if provinceName, exists := provinces[vote.Province]; exists {
			provinceVotes[provinceName]++
			continue
		}

		provinceCode := vote.Province
		if len(provinceCode) == 1 {
			provinceCode = "0" + provinceCode
		}

		if provinceName, exists := provinces[provinceCode]; exists {
			provinceVotes[provinceName]++
		}
	}

	resultsByProvince := make(map[string]EventImpactByProvince)

	for _, event := range events {
		provinceCode := event.Province
		if len(provinceCode) == 1 {
			provinceCode = "0" + provinceCode
		}
		provinceName, exists := provinces[provinceCode]
		if !exists {
			continue
		}

		var votesBefore int64
		beforeQuery := initializers.DB.Model(&models.MyVote{}).
			Where("candidate_id = ? AND created_at < ?", event.CandidateID, event.Date)

		if err := beforeQuery.Count(&votesBefore).Error; err != nil {
			fmt.Printf("Error counting votes before: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count votes before"})
			return
		}

		if votesBefore == 0 {
			votesBefore = 1 // Establecemos un valor mínimo
		}

		var votesAfter int64
		afterQuery := initializers.DB.Model(&models.MyVote{}).
			Where("candidate_id = ? AND province = ? AND created_at >= ? AND created_at <= ?",
				event.CandidateID, event.Province, event.Date, event.Date.Add(time.Hour*24*7))

		if err := afterQuery.Count(&votesAfter).Error; err != nil {
			fmt.Printf("Error counting votes after: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count votes after"})
			return
		}

		if votesAfter == 0 {
			votesAfter = 1 // Establecemos un valor mínimo
		}

		impactRatio := float64(votesAfter) / float64(votesBefore) * float64(event.Severity) / 10.0
		if math.IsNaN(impactRatio) || math.IsInf(impactRatio, 0) {
			impactRatio = 0.01 // Valor mínimo para evitar 0
		}

		if provincial, exists := resultsByProvince[provinceName]; exists {
			provincial.EventDetails = append(provincial.EventDetails, EventImpact{
				CandidateID:   event.CandidateID,
				Province:      provinceName,
				Date:          event.Date,
				Description:   event.Description,
				Severity:      event.Severity,
				VotesBefore:   votesBefore,
				VotesAfter:    votesAfter,
				ImpactRatio:   impactRatio,
			})
			provincial.TotalEvents++
			provincial.TotalVotesBefore += votesBefore
			provincial.TotalVotesAfter += votesAfter
			provincial.VoteImpactRatio = float64(provincial.TotalVotesAfter) / float64(provincial.TotalVotesBefore)
			if math.IsNaN(provincial.VoteImpactRatio) || math.IsInf(provincial.VoteImpactRatio, 0) {
				provincial.VoteImpactRatio = 0.01 // Valor mínimo para evitar 0
			}
			resultsByProvince[provinceName] = provincial
		} else {
			resultsByProvince[provinceName] = EventImpactByProvince{
				CandidateID:      event.CandidateID,
				Province:         provinceName,
				TotalEvents:      1,
				EventDetails:     []EventImpact{{
					CandidateID:   event.CandidateID,
					Province:      provinceName,
					Date:          event.Date,
					Description:   event.Description,
					Severity:      event.Severity,
					VotesBefore:   votesBefore,
					VotesAfter:    votesAfter,
					ImpactRatio:   impactRatio,
				}},
				TotalVotesBefore: votesBefore,
				TotalVotesAfter:  votesAfter,
				VoteImpactRatio:  impactRatio,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"candidate_id":      candidateID,
		"province_analysis": resultsByProvince,
	})
}