package model

import (
	"../db"
	_ "github.com/k0kubun/pp"
	_ "log"
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

func FindEntriesIndex(permit int, page int) ([]Entry, []Entry) {
	var entries []Entry
	var nextEntries []Entry
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&entries).Select("Title").Offset(page * permit).Limit(permit).Find(&nextEntries)
	return entries, nextEntries
}

func FindEntry(id string) (Entry, bool) {
	var entry Entry
	return entry, db.Dbmap.Find(&entry, id).RecordNotFound()
}
