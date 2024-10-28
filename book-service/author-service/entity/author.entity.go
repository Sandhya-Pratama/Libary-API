package entity

import "time"

type Author struct {
	ID        int64
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewAuthor(name, bio string) *Author {
	return &Author{
		Name:      name,
		Bio:       bio,
		CreatedAt: time.Now(),
	}
}

func UpdateAuthor(id int64, name, bio string) *Author {
	return &Author{
		ID:        id,
		Name:      name,
		Bio:       bio,
		UpdatedAt: time.Now(),
	}
}
