package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID
	Name             string
	Email            string
	Password         string
	Role             []UserRole
	CreatedDate      time.Time
	LastModifiedDate time.Time
}

type UserRole struct {
	Id               uuid.UUID
	Role             Role
	UserId           uuid.UUID
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
