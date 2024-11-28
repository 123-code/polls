package controllers

import (
	"fmt"
	"net/http"
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
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }


    userIDUint, err := strconv.ParseUint(request.UserID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "UserID must be a valid integer",
        })
        return
    }


    validate,err := validators.ValidateID(c,uint(userIDUint))
    if !validate {
        fmt.Println("validation error,UserID must be a valid integer")
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "UserID must be a valid integer",
        })
    } else{
    jwt,err := util.GenerateJWTs(c,string(userIDUint))
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(jwt)
    c.JSON(http.StatusAccepted, gin.H{
        "token": jwt,
    })
    }

}
    

  
