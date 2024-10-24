package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Roles     string     `json:"roles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewUser(Username, Email, Password, Roles string) *User {
	return &User{
		Username:  Username,
		Email:     Email,
		Password:  Password,
		Roles:     Roles,
		CreatedAt: time.Now(),
	}
}

func UpdateUser(Username, Email, Password, Roles string) *User {
	return &User{
		Username:  Username,
		Email:     Email,
		Password:  Password,
		Roles:     Roles,
		UpdatedAt: time.Now(),
	}

}
