package main

import (
	"./db"
	_ "./helper"
	"./models"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	_ "log"
	"strings"
	"testing"
)

func TestCreateEntryModel(t *testing.T) {

	title := "たべたいもの"
	content := `
# おなかがすいたんだおー。
## 食べたいものリストなんだお

- 寿司
- 天ぷら
- 焼き肉

まぁこんなもんかな。
`

	createEntry := model.Entry{
		Title:   title,
		Content: content,
		ThemeId: 1,
	}

	db.DbTestConnect()
	db.Dbmap.NewRecord(createEntry)
	db.Dbmap.Create(&createEntry)

	var entry model.Entry
	db.Dbmap.First(&entry)

	if entry.Title != title {
		t.Errorf("got %v want %v", entry.Title, title)
	}
}

func TestEntryGenerateHtmlFromMarkdown(t *testing.T) {
	db.DbTestConnect()

	var entry model.Entry
	db.Dbmap.Last(&entry)

	output := blackfriday.MarkdownCommon([]byte(entry.Content))
	html := string(output)
	r := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(r)

	h2 := doc.Find("h2").Text()
	heading := "食べたいものリストなんだお"

	if h2 != heading {
		t.Errorf("got %v want %v", h2, heading)
	}
}
