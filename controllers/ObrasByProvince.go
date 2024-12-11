package controllers

import (
	"fmt"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"

	"github.com/gin-gonic/gin"
)

func AnalyzeObrasByProvince(c *gin.Context) {
	var cedulas []models.Cedula
	var obras []models.Obra

	// Obtener todas las cedulas
	if err := initializers.DB.Find(&cedulas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user IDs",
		})
		return
	}

	// Obtener todas las obras
	if err := initializers.DB.Find(&obras).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve obras",
		})
		return
	}

	voterCount := make(map[string]int)
	obrasCount := make(map[string]int)

	// Contar votos por provincia
	for _, cedula := range cedulas {
		if len(cedula.UserID) < 2 {
			continue 
		}

		provinceCode := cedula.UserID[:2]
		if provinceName, exists := provinces[provinceCode]; exists {
			voterCount[provinceName]++
		}
	}

// Contar obras por provincia
// Contar obras por provincia
for _, obra := range obras {
    fmt.Println("Province from obra:", obra.Province) // Agrega esta línea para depurar
    if provinceName, exists := provinces[obra.Province]; exists {
        obrasCount[provinceName]++
    }
}

	// Calcular el ratio de obras a votos
	result := gin.H{}
	for province, count := range voterCount {
		obrasForProvince := obrasCount[province]
		ratio := 0.0
		if count > 0 {
			ratio = float64(obrasForProvince) / float64(count)
		}
		result[province] = gin.H{
			"votos": count,
			"obras": obrasForProvince,
			"ratio": ratio,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"obra_analysis": result,
	})
}