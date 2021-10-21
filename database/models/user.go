package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Email     string `gorm:"type:varchar(200);uniqueIndex;not null"`
	Password  string `gorm:"not null"`
}
