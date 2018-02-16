package routes

import (
	"time"
)

type signupCommand struct {
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Password    string    `json:"password" validate:"required,gte=8,lte=30"`
}

type loginCommand struct {
	email    string `validate:"nonzero"`
	password string `validate:"nonzero"`
}

type FieldError struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
}

type BadRequestError struct {
	Code    string
	Message string
	Details []FieldError
}
