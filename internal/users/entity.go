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

type Address struct {
	ID         int64     `json:"id"`
	Street     string    `json:"street"`
	Number     int       `json:"number"`
	Complement string    `json:"complement"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	PostalCode string    `json:"postal_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsAdmin   bool      `json:"-"`
	Address   Address   `json:"address"`
	AddressID int64     `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"-"`
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
