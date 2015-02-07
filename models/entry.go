package model

import (
	"time"
)

type Entry struct {
	Id        int64
	Title     string
	Content   string
	Theme     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
