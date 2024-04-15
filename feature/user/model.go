package user

import "github.com/google/uuid"

type UserContext struct {
	TenantId  uuid.UUID
	Id        uuid.UUID
	AccountId uuid.UUID
	Email     string
	LastName  string
	FirstName string
}

// Account
