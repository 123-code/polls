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
    candidateID := c.Param("id")

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

    c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully", "candidate": candidate})
}

func (vc *VoteController) GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	if err := initializers.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, candidates)
}
