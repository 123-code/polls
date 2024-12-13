package controllers

import (
    "net/http"
    "pollsbackend/initializers"
    "pollsbackend/models"
    "github.com/gin-gonic/gin"
)

type Candidate struct {
    ID           uint   `json:"id"`
    President    string `json:"president"`
    VicePresident string `json:"vicepresident"`
}

var candidates = []Candidate{
    {ID: 2, President: "Daniel Noboa", VicePresident: "Pinto"},
    {ID: 3, President: "Luisa Gonzalez", VicePresident: "Borja"},
    {ID: 4, President: "Leonidas Iza", VicePresident: "Molina"},
    {ID: 5, President: "Pedro Granja", VicePresident: "Silva"},
    {ID: 6, President: "Francesco Tabacchi", VicePresident: "Sacancela"},
    {ID: 7, President: "Enrique Gomez", VicePresident: "Inés Díaz"},
    {ID: 8, President: "Henry Cucalon", VicePresident: "Larrea"},
    {ID: 9, President: "Kronfle", VicePresident: "Passailaigue"},
    {ID: 10, President: "Carlos Rabascall", VicePresident: "Rivas"},
    {ID: 11, President: "Jorge Escala", VicePresident: "Terán"},
    {ID: 12, President: "Ivan Saquicela", VicePresident: "Coello"},
    {ID: 13, President: "Andrea Gonzalez", VicePresident: "Moncayo"},
    {ID: 14, President: "Juan Cueva", VicePresident: "Reyes"},
    {ID: 15, President: "Victor Araus", VicePresident: "Carrera"},
    {ID: 16, President: "Luis Tilleria", VicePresident: "Rosero"},
    {ID: 17, President: "Jimmy Jairala", VicePresident: "Vallecilla"},
}

type ObrasRatio struct {
    CandidateID    uint    `json:"candidate_id"`
    President      string  `json:"president"`
    VicePresident  string  `json:"vicepresident"`
    TotalVotes     int64   `json:"total_votes"`
    TotalObras     int64   `json:"total_obras"`
    Ratio          float64 `json:"ratio"`
}

func GetObrasRatio(c *gin.Context) {
    var results []ObrasRatio

    for _, candidate := range candidates {
        var voteCount int64
        var obrasCount int64
        var completedObrasCount int64
        var inProgressObrasCount int64

        if err := initializers.DB.Model(&models.MyVote{}).
            Where("candidate_id = ?", candidate.ID).
            Count(&voteCount).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count votes"})
            return
        }

        if err := initializers.DB.Model(&models.Obra{}).
            Where("candidate_id = ?", candidate.ID).
            Count(&obrasCount).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count obras"})
            return
        }

        if err := initializers.DB.Model(&models.Obra{}).
            Where("candidate_id = ? AND status = ?", candidate.ID, "completed").
            Count(&completedObrasCount).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count completed obras"})
            return
        }

        // Count in-progress obras
        if err := initializers.DB.Model(&models.Obra{}).
            Where("candidate_id = ? AND status = ?", candidate.ID, "in_progress").
            Count(&inProgressObrasCount).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count in-progress obras"})
            return
        }
        ratio := float64(0)
        if obrasCount > 0 {
   
            weightedObrasCount := float64(completedObrasCount) + (float64(inProgressObrasCount) * 0.5)
            ratio = float64(voteCount) / weightedObrasCount
        }

        results = append(results, ObrasRatio{
            CandidateID:    candidate.ID,
            President:      candidate.President,
            VicePresident:  candidate.VicePresident,
            TotalVotes:     voteCount,
            TotalObras:     obrasCount,
            Ratio:          ratio,
        })
    }

    c.JSON(http.StatusOK, results)
}

func InitializeObras(c *gin.Context) {

    obras := []models.Obra{
        {CandidateID: 2, Name: "Infrastructure Project A", Description: "Road construction", Status: "completed",Province:"01"},
        {CandidateID: 3, Name: "Healthcare Initiative B", Description: "Hospital renovation", Status: "in_progress",Province:"01"},
        {CandidateID: 4, Name: "Education Program C", Description: "New school construction", Status: "completed",Province:"01"},
        {CandidateID: 5, Name: "Agricultural Development D", Description: "Irrigation system installation", Status: "completed",Province:"01"},
        {CandidateID: 6, Name: "Public Transport E", Description: "Bus route expansion", Status: "in_progress",Province:"01"},
        {CandidateID: 7, Name: "Community Center F", Description: "Building a community center", Status: "completed",Province:"01"},
        {CandidateID: 8, Name: "Renewable Energy G", Description: "Solar panel installation", Status: "completed",Province:"01"},
        {CandidateID: 9, Name: "Waste Management H", Description: "Recycling program launch", Status: "in_progress",Province:"01"},
        {CandidateID: 10, Name: "Cultural Heritage I", Description: "Restoration of historical sites", Status: "completed",Province:"01"},
        {CandidateID: 11, Name: "Sports Facility J", Description: "Construction of a sports complex", Status: "in_progress",Province:"01"},
        {CandidateID: 12, Name: "Urban Development K", Description: "City park renovation", Status: "completed",Province:"01"},
        {CandidateID: 13, Name: "Technology Hub L", Description: "Establishment of a tech incubator", Status: "completed",Province:"01"},
        {CandidateID: 14, Name: "Disaster Relief M", Description: "Emergency response training program", Status: "in_progress",Province:"01"},
        {CandidateID: 15, Name: "Tourism Initiative N", Description: "Promotion of local tourism sites", Status: "completed",Province:"01"},
        {CandidateID: 16, Name: "Infrastructure Upgrade O", Description: "Road repair and maintenance", Status: "completed",Province:"01"},
        {CandidateID: 17, Name: "Healthcare Access P", Description: "Mobile health clinics program", Status: "in_progress",Province:"01"},
    }

    for _, obra := range obras {
        if err := initializers.DB.Create(&obra).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize obras"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Obras initialized successfully"})
}