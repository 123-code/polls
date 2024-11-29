package controllers

import (
	"fmt"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"

	"github.com/gin-gonic/gin"
)

func (vc *VoteController) GetAllVotes(c *gin.Context) {
    var votes []models.MyVote

    // Get sorting order from query parameters
    sortOrder := c.Query("sort")
    fmt.Println("Sort parameter received:", sortOrder) // Debugging log

    // Default to sorting by newest if no sort order is specified
    order := "created_at DESC" // Newest first
    if sortOrder == "oldest" {
        order = "created_at ASC" // Oldest first
    }

    // Query all votes and sort by the specified order
    if err := initializers.DB.Model(&models.MyVote{}).Order(order).Find(&votes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch votes"})
        return
    }

    c.JSON(http.StatusOK, votes)
}