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

    // Bind incoming JSON request
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Log received UserID
    fmt.Println("Received UserID:", request.UserID)

    // Validate that UserID is a valid integer string
    userIDUint, err := strconv.ParseUint(request.UserID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserID must be a valid integer"})
        return
    }

    // Perform validation using the parsed uint value
    validate, err := validators.ValidateID(c, uint(userIDUint))
    if !validate {
        fmt.Println("Validation failed for UserID:", request.UserID)
        fmt.Println("Error:", err) // Log any error returned by ValidateID
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserID must be a valid integer"})
        return
    }

    // Insert the user ID into the cedulas table
    cedula := models.Cedula{UserID: request.UserID} // Store UserID as string
    if err := initializers.DB.Create(&cedula).Error; err != nil {
        fmt.Println("Error inserting into cedulas:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user ID"})
        return
    }

    // Generate JWT token
    jwt, err := util.GenerateJWTs(c, request.UserID) // Pass UserID as string
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    fmt.Println(jwt)
    c.JSON(http.StatusAccepted, gin.H{"token": jwt})
}