package models

import (
    "gorm.io/gorm"
    "time"
)

type Event struct {
    gorm.Model
    CandidateID  uint      `json:"candidate_id"`
    Province     string    `json:"province"`
    Date         time.Time `json:"date"`
    Description  string    `json:"description"`
    Severity     int       `json:"severity"`
}