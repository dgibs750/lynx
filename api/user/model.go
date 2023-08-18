package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"firstName" validate:"required,alpha_space"`
	LastName       string    `json:"lastName" validate:"required,alpha_space"`
	Email          string    `json:"email" validate:"required,email"`
	MasterPassword string    `json:"-"`
	CreatedAt      string    `json:"-"`
}

type NewUser struct {
	FirstName      string `json:"firstName" validate:"required,alpha_space"`
	LastName       string `json:"lastName" validate:"required,alpha_space"`
	Email          string `json:"email" validate:"required,email"`
	MasterPassword string `json:"masterPassword" validate:"required,min=6,pwd"`
}

type UpdateUserData struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"firstName" validate:"required,alpha_space"`
	LastName       string    `json:"lastName" validate:"required,alpha_space"`
	Email          string    `json:"email" validate:"required,email"`
	MasterPassword string    `json:"masterPassword" validate:"required,min=6,pwd"`
	CreatedAt      string    `json:"-"`
}
