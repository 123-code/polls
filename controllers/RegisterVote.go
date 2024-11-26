package controllers

import (
	"fmt"
	"net/http"
	"pollsbackend/models"

	//"pollsbackend/util"
	"github.com/gin-gonic/gin"
	"pollsbackend/initializers"
	"gorm.io/gorm"
)

type VoteController struct {

}

type vote struct{
	gorm.Model
	CandidateID uint   `json:"candidate_id"`
	IPAddress   string `json:"ip_address"`
}

func NewVoteController(db *gorm.DB) *VoteController {
	return &VoteController{}
}

func (vc *VoteController) RegisterCandidate(c *gin.Context) {
	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if initializers.DB == nil {
		fmt.Println("DB is nil, cannot create candidate")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not established"})
		return
	}

	result := initializers.DB.Create(&candidate)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}
func (vc *VoteController) CastVote(c *gin.Context) {
    var vote models.MyVote  
    candidateID := c.Param("id")
    ipAddress := c.ClientIP()
    var totalVotesCount int64
    result := initializers.DB.Model(&models.MyVote{}).Where("ip_address = ?", ipAddress).Count(&totalVotesCount)
    
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking vote count"})
        return
    }

    // Check if IP has already reached vote limit
    if totalVotesCount >= 5 {
        c.JSON(http.StatusForbidden, gin.H{
            "error": "Maximum vote limit (5 votes) reached from this IP address",
            "votes_made": totalVotesCount,
        })
        return
    }

    var existingVote models.MyVote
    duplicateVoteCheck := initializers.DB.Where("candidate_id = ? AND ip_address = ?", candidateID, ipAddress).First(&existingVote)
    
    if duplicateVoteCheck.Error == nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "You have already voted for this candidate"})
        return
    }

    var candidate models.Candidate
    if err := initializers.DB.First(&candidate, candidateID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
        return
    }

    candidate.Votes++
    if err := initializers.DB.Save(&candidate).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote count"})
        return
    }

    // Record the vote
    vote = models.MyVote{
        CandidateID: candidate.ID,
        IPAddress:   ipAddress,
    }
    
    if err := initializers.DB.Create(&vote).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record vote"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Vote cast successfully",
        "candidate": candidate,
        "votes_remaining": 5 - (totalVotesCount + 1),
    })
}

func (vc *VoteController) GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	if err := initializers.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, candidates)
}
