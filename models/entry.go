package model

import (
	"time"
)

type Entry struct {
	Id         int64
	Title      string
	Content    string
	Basename   string
	AuthoredOn time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
