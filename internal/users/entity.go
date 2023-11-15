package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

func New(name string, age int, email string, password string, isAdmin bool, addressID int64) (*User, error) {
	now := time.Now()
	u := User{
		Name:      name,
		Age:       age,
		Email:     email,
		IsAdmin:   isAdmin,
		AddressID: addressID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := u.SetPassword(password)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

type User struct {
	ID        int64
	Name      string
	Age       int
	Email     string
	Password  string
	IsAdmin   bool
	AddressID int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return errors.New("passwords is required and can't be empty")
	}

	if len(password) < 6 {
		return errors.New("passwords must have at least 6 characters long")
	}

	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	return nil
}
