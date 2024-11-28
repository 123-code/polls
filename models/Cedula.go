package models

import (
	"gorm.io/gorm"
)

type Cedula struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"unique;not null"`
}