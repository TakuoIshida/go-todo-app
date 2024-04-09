package tenant

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	CreatedAt       time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt       time.Time `gorm:"type:timestamp;not null;default:now()"`
	CreateUserId    uuid.UUID `gorm:"type:char(36);not null"`
	UpdateUserId    uuid.UUID `gorm:"type:char(36);not null"`
	TenantId        uuid.UUID `gorm:"type:char(36);not null;index"`
	Id              uuid.UUID `gorm:"type:char(36);not null;primary_key;index"`
	Title           string    `gorm:"type:varchar(255);not null"`
	Description     string    `gorm:"type:text"`
	IsDeleted       bool      `gorm:"type:boolean;default:false"`
	AttachmentFiles []AttachmentFile
}

type AttachmentFile struct {
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:now()"`
	CreateUserId uuid.UUID `gorm:"type:char(36);not null"`
	UpdateUserId uuid.UUID `gorm:"type:char(36);not null"`
	TenantId     uuid.UUID `gorm:"type:char(36);not null;index"`
	Id           uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Size         int64     `gorm:"type:bigint;not null"`
	ContentType  string    `gorm:"type:varchar(255);not null"`
	IsDeleted    bool      `gorm:"type:boolean;default:false"`
	TodoId       uuid.UUID `gorm:"type:char(36);not null;constraint:OnDelete:CASCADE;"`
}
