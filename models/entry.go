package model

import (
	"time"
)

type Entry struct {
	Id        int
	Title     string
	Content   string
	ThemeId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
