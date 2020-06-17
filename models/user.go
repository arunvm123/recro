package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID          int            `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" gorm:"not null"`
	Email       string         `json:"email" gorm:"unique"`
	PhoneNumber string         `json:"phone_number" gorm:"index:phone"`
	Meta        postgres.Jsonb `json:"meta"`
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Save(db *gorm.DB) error {
	return db.Save(&u).Error
}
