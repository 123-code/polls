package controllers

import (
	"fmt"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"pollsbackend/util"
	"pollsbackend/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserIDRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

func EnterUser(c *gin.Context) {
    var request UserIDRequest


    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

   
    fmt.Println("Received UserID:", request.UserID)


    userIDUint, err := strconv.ParseUint(request.UserID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserID must be a valid integer"})
        return
    }


    validate, err := validators.ValidateID(c, uint(userIDUint))
    if !validate {
        fmt.Println("Validation failed for UserID:", request.UserID)
        fmt.Println("Error:", err) 
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserID must be a valid integer"})
        return
    }


    cedula := models.Cedula{UserID: request.UserID} 
    if err := initializers.DB.Create(&cedula).Error; err != nil {
        fmt.Println("Error inserting into cedulas:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user ID"})
        return
    }


    jwt, err := util.GenerateJWTs(c, request.UserID) 
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    fmt.Println(jwt)
    c.JSON(http.StatusAccepted, gin.H{"token": jwt})
}