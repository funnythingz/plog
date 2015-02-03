package main

import (
	"./db"
	"./models"
	"testing"
)

func TestEntryModel(t *testing.T) {

	title := "うひょひょひょー"
	content := "ひょひょひょのひょーーーーーー"
	basename := "post-1"

	createEntry := model.Entry{
		Title:    title,
		Content:  content,
		Basename: basename,
	}

	dbmap.DbTestConnect()
	dbmap.Dbmap.NewRecord(createEntry)
	dbmap.Dbmap.Create(&createEntry)

	var entry model.Entry
	dbmap.Dbmap.First(&entry)

	if entry.Title != title {
		t.Errorf("got %v want %v", entry.Title, title)
	}

	dbmap.Dbmap.Delete(&entry)
}
