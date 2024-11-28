package controllers

import (
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	//"strconv"
	//"strings"
	"github.com/gin-gonic/gin"
)

// Province codes and their names
var provinces = map[string]string{
	"01": "Azuay",
	"02": "Bolivar",
	"03": "Cañar",
	"04": "Carchi",
	"05": "Cotopaxi",
	"06": "Chimborazo",
	"07": "El Oro",
	"08": "Esmeraldas",
	"09": "Guayas",
	"10": "Imbabura",
	"11": "Loja",
	"12": "Los Rios",
	"13": "Manabi",
	"14": "Morona Santiago",
	"15": "Napo",
	"16": "Pastaza",
	"17": "Pichincha",
	"18": "Tungurahua",
	"19": "Zamora Chinchipe",
	"20": "Galapagos",
	"21": "Sucumbios",
	"22": "Orellana",
	"23": "Santo Domingo de los Tsachilas",
	"24": "Santa Elena",
}

func AnalyzeVotersByProvince(c *gin.Context) {
	var cedulas []models.Cedula

	// Fetch all user IDs from the database
	if err := initializers.DB.Find(&cedulas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user IDs",
		})
		return
	}

	// Initialize a map to count voters per province
	voterCount := make(map[string]int)

	// Process each user ID
	for _, cedula := range cedulas {
		if len(cedula.UserID) < 2 {
			continue // Skip invalid IDs
		}

		provinceCode := cedula.UserID[:2] // Extract first two digits
		if provinceName, exists := provinces[provinceCode]; exists {
			voterCount[provinceName]++
		}
	}

	// Prepare the result
	result := gin.H{}
	for province, count := range voterCount {
		result[province] = count
	}

	// Return the analysis
	c.JSON(http.StatusOK, gin.H{
		"voter_analysis": result,
	})
}
