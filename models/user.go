package models

import (
	"encoding/json"

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

type OauthSignupArgs struct {
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	PhoneNumber  string       `json:"phone_number"`
	ProviderData ProviderData `json:"provider_data"`
}

type ProviderData struct {
	Github   interface{} `json:"github"`
	LinkedIn interface{} `json:"linkedIn"`
}

type UserInfo struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber *string         `json:"phone_number"`
	Meta        *postgres.Jsonb `json:"meta"`
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

func UserOauthSignup(db *gorm.DB, args *OauthSignupArgs, provider string) (*User, error) {
	user := User{
		Email:       args.Email,
		Name:        args.Name,
		PhoneNumber: &args.PhoneNumber,
	}

	rawJson, err := json.Marshal(args.ProviderData)
	if err != nil {
		log.WithFields(log.Fields{
			"func":     "UserOauthSignup",
			"subFunc":  "json.Marshal",
			"email":    args.Email,
			"provider": provider,
		}).Error(err)
		return nil, err
	}

	user.Meta = &postgres.Jsonb{
		RawMessage: json.RawMessage(rawJson),
	}

	err = user.Create(db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":     "UserOauthSignup",
			"subFunc":  "user.Create",
			"email":    args.Email,
			"provider": provider,
		}).Error(err)
		return nil, err
	}

	return &user, nil
}

func UpdateProviderDetails(db *gorm.DB, email, provider string, providerData interface{}) (*User, error) {
	user, err := GetUserFromEmail(db, email)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UpdateProviderDetails",
			"subFunc": "GetUserFromEmail",
			"email":   email,
		}).Error(err)
		return nil, err
	}

	var data ProviderData
	err = json.Unmarshal(user.Meta.RawMessage, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UpdateProviderDetails",
			"subFunc": "json.Unmarshal",
			"email":   email,
		}).Error(err)
		return nil, err
	}

	switch provider {
	case "github":
		data.Github = providerData
	}

	dataBytes, err := json.Marshal(&data)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UpdateProviderDetails",
			"subFunc": "json.Unmarshal",
			"email":   email,
		}).Error(err)
		return nil, err
	}

	user.Meta.RawMessage = json.RawMessage(dataBytes)

	err = user.Save(db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "UpdateProviderDetails",
			"subFunc": "user.Save",
			"email":   email,
		}).Error(err)
		return nil, err
	}

	return user, nil
}

func GetAllUsers(db *gorm.DB) (*[]UserInfo, error) {
	var users []UserInfo

	err := db.Table("users").Find(&users).Error
	if err != nil {
		log.WithFields(log.Fields{
			"func": "GetAllUsers",
			"info": "retrieving details of all users",
		}).Error(err)
		return nil, err
	}

	return &users, nil
}
