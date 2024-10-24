package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewUser(username, email, password, role string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      role,
		CreatedAt: time.Now(),
	}
}

// Public Register
func Register(username, email, password, role string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func UpdateUser(id int64, username, email, password, role string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      role,
		UpdatedAt: time.Now(),
	}

}
