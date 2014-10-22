package model

import (
	"errors"
	"strings"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" sql:"not null;unique;size:255"`
	Email string `json:"email"`
}

func GetAllUsers() (users []User) {
	users = make([]User, 0)
	db.Find(&users)
	return
}

func GetUser(id int64) *User {
	var user User
	err := db.Find(&user, id).Error
	if err != nil {
		return nil
	} else {
		return &user
	}
}

func DeleteUser(user *User) {
	db.Delete(user)
}

func CreateNewUser(user *User) error {
	if err := user.IsValid(); err != nil {
		return err
	}
	result := db.Create(user)
	return result.Error
}

func (u *User) IsValid() error {
	if len(strings.TrimSpace(u.Name)) == 0 {
		return errors.New("Name can't be blank")
	}
	return nil
}
