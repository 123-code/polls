package models

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Votes int    `json:"votes" gorm:"default:0"`
}