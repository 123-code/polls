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

    if len(obra.Province) == 1 {
        obra.Province = "0" + obra.Province
    }

    if _, exists := provinces[obra.Province]; !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province code"})
        return
    }

    if err := initializers.DB.Create(&obra).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create obra"})
        return
    }

    c.JSON(http.StatusOK, obra)
}