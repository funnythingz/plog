package model

import (
	"../db"
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

func FindEntriesIndex(permit int, page int) ([]Entry, bool, bool) {
	var entries []Entry
	var nextEntries []Entry
	current := db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&entries).Select("Title")
	endpoint := current.Offset(page * permit).Find(&nextEntries).RecordNotFound()
	return entries, current.RecordNotFound(), endpoint
}

func FindEntry(id string) (Entry, bool) {
	var entry Entry
	return entry, db.Dbmap.Find(&entry, id).RecordNotFound()
}
