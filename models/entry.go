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

func FindEntriesIndex() ([]Entry, bool) {
	var entries []Entry
	return entries, db.Dbmap.Order("id desc").Find(&entries).Select("Title").RecordNotFound()
}

func FindEntry(id string) (Entry, bool) {
	var entry Entry
	return entry, db.Dbmap.Find(&entry, id).RecordNotFound()
}
