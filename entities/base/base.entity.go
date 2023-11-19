package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ID        uint64 `gorm:"primaryKey"`
type BaseEntity struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
