package models

import (
	"errors"
	"log"
	"time"

	"github.com/hillfolk/go-ddd-start/db"
	"github.com/hillfolk/go-ddd-start/forms"
	"github.com/satori/go.uuid"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}

	if err := db.Save(user); err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return &user, nil
}

func (h User) GetByID(id string) (*User) {
	db := db.GetDB()
	var user *User
	db.First(user,id)
	return user
}