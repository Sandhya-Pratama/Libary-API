package entity

import "time"

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewCategory(name string) *Category {
	return &Category{
		Name:      name,
		CreatedAt: time.Now(),
	}
}

func UpdateCategory(id int64, name string) *Category {
	return &Category{
		ID:        id,
		Name:      name,
		UpdatedAt: time.Now(),
	}
}
