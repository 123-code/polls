package initializers

import "pollsbackend/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.Candidate{})

}