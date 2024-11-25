package controllers

import (
	"net/http"
	"pollsbackend/models"
	//"pollsbackend/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VoteController struct {
	DB *gorm.DB
}

func NewVoteController(db *gorm.DB) *VoteController {
	return &VoteController{DB: db}
}

func (vc *VoteController) RegisterCandidate(c *gin.Context) {
	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := vc.DB.Create(&candidate)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

func (vc *VoteController) CastVote(c *gin.Context) {
	candidateID := c.Param("id")

	var candidate models.Candidate
	if err := vc.DB.First(&candidate, candidateID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}

	candidate.Votes++
	if err := vc.DB.Save(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

func (vc *VoteController) GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	if err := vc.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, candidates)
}
