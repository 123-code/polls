package controllers

import (
    "net/http"
    "pollsbackend/initializers"
    "pollsbackend/models"
    "github.com/gin-gonic/gin"
)

func CreateObra(c *gin.Context) {
    var obra models.Obra

    if err := c.ShouldBindJSON(&obra); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := initializers.DB.Create(&obra).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create obra"})
        return
    }

    c.JSON(http.StatusOK, obra)
}