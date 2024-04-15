package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id           uuid.UUID `json:"id"`
	TenantId     uuid.UUID `json:"tenantId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	IsDeleted    bool      `json:"isDeleted"`
	CreatedAt    time.Time `json:"createdAt"`
	CreateUserId uuid.UUID `json:"createUserId"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdateUserId uuid.UUID `json:"updateUserId"`
	// AttachmentFiles []AttachmentFile
}

func New(title string, description string, userId uuid.UUID, tenantId uuid.UUID) (*Todo, error) {
	uuid := uuid.New()
	now := time.Now()
	todo := Todo{
		Id:           uuid,
		TenantId:     tenantId,
		Title:        title,
		Description:  description,
		IsDeleted:    false,
		CreatedAt:    now,
		CreateUserId: userId,
		UpdatedAt:    now,
		UpdateUserId: userId,
	}

	if err := validate(&todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func validate(t *Todo) error {
	MAX_LENGTH := 255

	if t.Title == "" {
		return errors.New("title is required")
	}
	if len(t.Title) > MAX_LENGTH {
		return errors.New("title is too long")
	}
	if len(t.Description) > MAX_LENGTH {
		return errors.New("description is too long")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}

	return nil
}
