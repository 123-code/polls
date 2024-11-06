package controllers

import (
	"net/http"
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
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    // Convert UserID string to uint for validation
    userIDUint, err := strconv.ParseUint(request.UserID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "UserID must be a valid integer",
        })
        return
    }

    // Call the validation function
    validate,err := validators.ValidateID(uint(userIDUint))
    if validate != true{
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "UserID must be a valid integer",
        })
    }

}
    

  
