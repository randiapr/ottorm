package entity

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
}
