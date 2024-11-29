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


    sortOrder := c.Query("sort")
    fmt.Println("Sort parameter received:", sortOrder) 


    order := "created_at DESC" 
    if sortOrder == "oldest" {
        order = "created_at ASC" 
    }


    if err := initializers.DB.Model(&models.MyVote{}).Order(order).Find(&votes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch votes"})
        return
    }

    c.JSON(http.StatusOK, votes)
}