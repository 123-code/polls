package models

import (
    "gorm.io/gorm"
)

type Obra struct {
    gorm.Model
    CandidateID uint   `json:"candidate_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Status      string `json:"status"`
    Province    string `json:"province"`
}