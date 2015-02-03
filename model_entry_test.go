package main

import (
	"./db"
	"./models"
	"log"
	"testing"
)

func TestEntryModel(t *testing.T) {

	title := "うひょひょひょー"
	content := "ひょひょひょのひょーーーーーー"
	basename := "post-1"

	entry := model.Entry{
		Title:    title,
		Content:  content,
		Basename: basename,
	}

	dbmap.DbTestConnect()
	dbmap.Dbmap.NewRecord(entry)
	dbmap.Dbmap.Create(&entry)
	firstEntry := dbmap.Dbmap.First(&entry)

	log.Println(firstEntry)
	if entry.Title != title {
		t.Errorf("got %v want %v", entry.Title, title)
	}
}
