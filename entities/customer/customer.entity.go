package entities

import (
	entities "example.com/goproject9/entities/base"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	BaseProperties entities.BaseEntity `gorm:"embedded"`
	Firstname      string              `gorm:"not null;size:50"`
	Lastname       string              `gorm:"not null;size:50"`
	Email          string              `gorm:"unique;not null;size:50"`
	Age            uint8               `gorm:"default:0;size:99"`
	Gender         string              `gorm:"default:'';size:10"`
}
