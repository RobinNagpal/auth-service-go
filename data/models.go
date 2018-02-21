package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        string    `gorm:"type:varchar(255);unique_index;not null"`
	DateOfBirth  time.Time `gorm:"type:not null"`
	Password     string    `gorm:"type:varchar(255):not null"`
	TwoFAEnabled *bool
}

func (u User) ValidateModel() error {
	return nil
}

type ValidateModel interface {
	validate() error
}
