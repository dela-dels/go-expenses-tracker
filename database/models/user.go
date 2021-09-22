package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
