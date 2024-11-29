package controllers

import (
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"github.com/gin-gonic/gin"
)

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

	if err := initializers.DB.Find(&cedulas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user IDs",
		})
		return
	}

	voterCount := make(map[string]int)

	for _, cedula := range cedulas {
		if len(cedula.UserID) < 2 {
			continue 
		}

		provinceCode := cedula.UserID[:2]
		if provinceName, exists := provinces[provinceCode]; exists {
			voterCount[provinceName]++
		}
	}

	result := gin.H{}
	for province, count := range voterCount {
		result[province] = count
	}
	c.JSON(http.StatusOK, gin.H{
		"voter_analysis": result,
	})
}
