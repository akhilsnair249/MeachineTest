package main

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}

type Article struct {
	BaseModel   `valid:"-"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Name        string    `gorm:"type:varchar(32);not null"`
	EmailID     string    `gorm:"type:varchar(32);not null"`
	ServerTime  time.Time `gorm:"not null"`
}

type TokenData struct {
	UserCode string
	jwt.StandardClaims
}
