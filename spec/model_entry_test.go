package main

import (
	"../models"
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

	if entry.Title != title {
		t.Errorf("got %v want %v", entry.Title, title)
	}
}
