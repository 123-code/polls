package initializers

import "pollsbackend/models"

func SyncDatabase() {
    DB.AutoMigrate(&models.Candidate{})
    DB.AutoMigrate(&models.MyVote{})
    DB.AutoMigrate(&models.Cedula{})
    DB.AutoMigrate(&models.Obra{})
    DB.AutoMigrate(&models.Event{})
}