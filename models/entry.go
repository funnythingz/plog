package model

import (
	"github.com/funnythingz/plog/db"
	_ "github.com/k0kubun/pp"
	"time"
)

type Entry struct {
	Id        int
	Title     string
	Content   string
	Theme     string
	CreatedAt time.Time
	UpdatedAt time.Time

	Comments []Comment
}

func FindEntriesIndex(permit int, page int) ([]Entry, []Entry) {
	var entries []Entry
	var nextEntries []Entry
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&entries).Select("Title").Offset(page * permit).Limit(permit).Find(&nextEntries)
	return entries, nextEntries
}

func FindEntry(id string) (Entry, bool) {
	var entry Entry
	var comments []Comment

	isNotFound := db.Dbmap.Find(&entry, id).RecordNotFound()

	db.Dbmap.Model(&comments).Order("id desc").Model(&entry).Related(&comments)
	db.Dbmap.Model(&entry).Association("Comments").Append(&comments)

	return entry, isNotFound
}
