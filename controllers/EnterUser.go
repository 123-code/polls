package controllers

import (
	"net/http"
	"pollsbackend/validators"

	"github.com/gin-gonic/gin"
)

type UserIDRequest struct {
    UserID uint `json:"user_id" binding:"required"`
}


func EnterUser(c *gin.Context) {
    var request UserIDRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }
    isValid, err := validators.ValidateID(request.UserID)
    if err != nil || !isValid {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid ID format or unauthorized ID",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "User validated successfully",
    })
    }
    

  
