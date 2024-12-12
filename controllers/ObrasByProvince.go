package controllers

import (
	"fmt"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AnalyzeObrasByProvince(c *gin.Context) {
    var cedulas []models.Cedula
    var obras []models.Obra

    // Retrieve all cedulas 
    if err := initializers.DB.Find(&cedulas).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to retrieve user IDs",
        })
        return
    }

    // Retrieve all obras
    if err := initializers.DB.Find(&obras).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to retrieve obras",
        })
        return
    }

    voterCount := make(map[string]int)
    obrasCount := make(map[string]int)

    // Count votes by province
    for _, cedula := range cedulas {
        if len(cedula.UserID) < 2 {
            continue
        }
        
        provinceCode := cedula.UserID[:2]
        if provinceName, exists := provinces[provinceCode]; exists {
            voterCount[provinceName]++
        }
    }

    // Count obras by province
    for _, obra := range obras {
        fmt.Printf("Obra Province Raw: '%s'\n", obra.Province)
        
        // Try matching with and without trimming whitespace
        provinceName := provinces[obra.Province]
        if provinceName != "" {
            obrasCount[provinceName]++
        } else {
            // Try trimming whitespace
            trimmedProvince := strings.TrimSpace(obra.Province)
            for code, name := range provinces {
                if trimmedProvince == code {
                    obrasCount[name]++
                    break
                }
            }
        }
    }

    // Calculate the ratio of obras to votes
    result := gin.H{}
    for province, _ := range provinces {
        provinceName := provinces[province]
        obrasForProvince := obrasCount[provinceName]
        votesForProvince := voterCount[provinceName]
        
        ratio := 0.0
        if votesForProvince > 0 {
            ratio = float64(obrasForProvince) / float64(votesForProvince)
        }
        
        result[provinceName] = gin.H{
            "votos": votesForProvince,
            "obras": obrasForProvince,
            "ratio": ratio,
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "obra_analysis": result,
    })
}