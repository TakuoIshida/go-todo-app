package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	id           uuid.UUID
	tenantId     uuid.UUID
	title        string
	description  string
	isDeleted    bool
	updatedAt    time.Time
	updateUserId uuid.UUID
	// AttachmentFiles []AttachmentFile
}

func New(title string, description string, userId uuid.UUID, tenantId uuid.UUID) (*Todo, error) {
	uuid := uuid.New()
	todo := Todo{
		id:          uuid,
		tenantId:    tenantId,
		title:       title,
		description: description,
		isDeleted:   false,
	}

	if err := validate(&todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func Restore(t *Todo) *Todo {
	return &Todo{
		id:           t.id,
		tenantId:     t.tenantId,
		title:        t.title,
		description:  t.description,
		isDeleted:    t.isDeleted,
		updatedAt:    t.updatedAt,
		updateUserId: t.updateUserId,
	}
}

func validate(t *Todo) error {
	MAX_LENGTH := 255

	if t.title == "" {
		return errors.New("title is required")
	}
	if len(t.title) > MAX_LENGTH {
		return errors.New("title is too long")
	}
	if len(t.description) > MAX_LENGTH {
		return errors.New("description is too long")
	}
	if t.description == "" {
		return errors.New("description is required")
	}

	return nil
}
