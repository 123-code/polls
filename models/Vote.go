package models

import (
    "gorm.io/gorm"
)

type MyVote struct {
    gorm.Model
    CandidateID uint   `json:"candidate_id"`
    IPAddress   string `json:"ip_address"`
    Province    string `json:"province"` 
}