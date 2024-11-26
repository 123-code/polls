package initializers

import "pollsbackend/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.Candidate{})
	DB.AutoMigrate(&models.MyVote{})
}