package common

import (
	"time"

	"github.com/google/uuid"
)

type Example struct {
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:now()"`
	CreateUserId uuid.UUID `gorm:"type:char(36);not null"`
	UpdateUserId uuid.UUID `gorm:"type:char(36);not null"`
	TenantId     uuid.UUID `gorm:"type:char(36);not null;index"`
	Id           uuid.UUID `gorm:"type:char(36);not null;primary_key;index"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"type:varchar(255);not null"`
	IsDeleted    bool      `gorm:"type:boolean;default:false"`
}
