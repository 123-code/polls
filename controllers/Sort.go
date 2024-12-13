package controllers

import (
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"github.com/gin-gonic/gin"
)


func (vc *VoteController) SortVotes(c *gin.Context) {
    var votes []models.MyVote


    candidateID := c.Query("candidate_id")
    sortOrder := c.Query("sort")


    order := "created_at DESC" 
    if sortOrder == "oldest" {
        order = "created_at ASC"
    }

    query := initializers.DB.Model(&models.MyVote{}).Order(order)

   
    if candidateID != "" {
        query = query.Where("candidate_id = ?", candidateID)
    }


    if err := query.Find(&votes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch votes"})
        return
    }

    c.JSON(http.StatusOK, votes)
}