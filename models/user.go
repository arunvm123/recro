package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

type User struct {
	ID          int             `json:"id" gorm:"primary_key"`
	Name        string          `json:"name" gorm:"not null"`
	Email       string          `json:"email" gorm:"unique"`
	PhoneNumber *string         `json:"phone_number" gorm:"index:phone"`
	Password    string          `json:"password"`
	Meta        *postgres.Jsonb `json:"meta"`
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Save(db *gorm.DB) error {
	return db.Save(&u).Error
}

type SignUpArgs struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type LoginArgs struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func CheckIfUserExists(db *gorm.DB, email string) bool {
	var count int
	err := db.Table("users").Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "CheckIfUserExists",
			"info":  "checking if user with specified email exitst",
			"email": email,
		}).Error(err)
		return true
	}

	if count > 0 {
		return true
	}

	return false
}

func UserSignup(db *gorm.DB, args *SignUpArgs) error {
	var user User

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UserSignup",
			"subFunc": "bcrypt.GenerateFromPassword",
			"email":   args.Email,
		}).Error(err)
		return err
	}

	user.Password = string(passwordHash)
	user.Email = args.Email
	user.Name = args.Name
	user.PhoneNumber = &args.Password

	err = user.Create(db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UserSignup",
			"subFunc": "user.Create",
			"email":   args.Email,
		}).Error(err)
		return err
	}

	return nil
}

// GetUserFromEmail returns user details from the given email id
func GetUserFromEmail(db *gorm.DB, email string) (*User, error) {
	var user User

	err := db.Find(&user, "email = ?", email).Error
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "GetUserFromEmail",
			"info":  "retrieving user info from email",
			"email": email,
		}).Error(err)
		return nil, err
	}

	return &user, nil
}
