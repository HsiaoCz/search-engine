package types

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Loginform struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type User struct {
	ID       string     `bson:"_id" json:"id"`
	Email    string     `bson:"email" json:"email"`
	Password string     `bson:"password" json:"-"`
	IsAdmin  bool       `bson:"isAdmin" json:"isAdmin"`
	CreateAt *time.Time `bson:"createAt"`
	UpdateAt time.Time  `bson:"updateAt"`
}

func (u *User) CreateAdmin() (*User, error) {
	user := &User{
		Email:    "gg@gg.com",
		IsAdmin:  true,
		Password: "passwd",
	}
	// Has Password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.New("error creating password")
	}
	user.Password = string(password)
	// Create User in DB
	return user, nil
}
