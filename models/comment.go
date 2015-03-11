package model

import (
	_ "../db"
	_ "github.com/k0kubun/pp"
	_ "log"
	"time"
)

type Comment struct {
	Id        int
	EntryId   int
	Content   string
	CreatedAt time.Time
}
