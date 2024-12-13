package controllers
// recuperamos obras y votos para el candidato, los comparamos en una formula que calcula el ratio obras/votos. esto por cada provincia
import (
	"fmt"
	"net/http"
	"pollsbackend/initializers"
	"pollsbackend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AnalyzeObrasByProvince(c *gin.Context) {
  // Obtener el `candidate_id` desde el request
  candidateIDStr := c.Query("candidate_id")
  if candidateIDStr == "" {
      c.JSON(http.StatusBadRequest, gin.H{"error": "candidate_id is required"})
      return
  }

  candidateID, err := strconv.ParseUint(candidateIDStr, 10, 64)
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate_id"})
      return
  }


  var events []models.Event
  if err := initializers.DB.Where("candidate_id = ?", candidateID).Find(&events).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
      return
  }

  // Analizar votos por provincia
  provinceVotes := make(map[string]int)

  for _, provinceName := range provinces {
      provinceVotes[provinceName] = 0
  }

  var votes []models.MyVote
  if err := initializers.DB.Where("candidate_id = ?", candidateID).Find(&votes).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve votes"})
      return
  }

  for _, vote := range votes {
      fmt.Printf("Vote Province: %s\n", vote.Province)

      if provinceName, exists := provinces[vote.Province]; exists {
          provinceVotes[provinceName]++
          continue
      }

      // Manejo de códigos de provincia con un dígito
      provinceCode := vote.Province
      if len(provinceCode) == 1 {
          provinceCode = "0" + provinceCode
      }

      if provinceName, exists := provinces[provinceCode]; exists {
          provinceVotes[provinceName]++
      }
  }

  // Construir el resultado
  result := make(map[string]gin.H)

  for _, event := range events {
      provinceName := provinces[event.Province]

      votesBefore := provinceVotes[provinceName] // Votos antes del evento
      votesAfter := 0                            // Puedes calcular esto según tu lógica específica

      result[provinceName] = gin.H{
          "votes_before": votesBefore,
          "votes_after":  votesAfter,
          "impact_ratio": 0.0, // Calcular según tu lógica
      }
  }

  c.JSON(http.StatusOK, gin.H{
      "candidate_id":    candidateID,
      "province_votes":  provinceVotes,
      "province_events": result,
  })
}