package controllers

import (
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"github.com/gin-gonic/gin"
)


func (vc *VoteController) SortVotes(c *gin.Context) {
    var votes []models.MyVote

    // Get query parameters
    candidateID := c.Query("candidate_id")
    sortOrder := c.Query("sort")

    // Default to sorting by newest if no sort order is specified
    order := "created_at DESC" // Newest first
    if sortOrder == "oldest" {
        order = "created_at ASC" // Oldest first
    }

    query := initializers.DB.Model(&models.MyVote{}).Order(order)

    // Apply candidate_id filter if specified
    if candidateID != "" {
        query = query.Where("candidate_id = ?", candidateID)
    }

    // Execute query to fetch votes
    if err := query.Find(&votes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch votes"})
        return
    }

    c.JSON(http.StatusOK, votes)
}