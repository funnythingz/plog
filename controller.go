package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	_ "github.com/goji/param"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	_ "log"
	"net/http"
	"strconv"
)

func top(c web.C, w http.ResponseWriter, r *http.Request) {

	Entries, entriesNotFound := model.FindEntriesIndex()

	if entriesNotFound {
		NotFound(w, r)
		return
	}

	tpl, _ := ace.Load("views/layouts/layout", "views/top", nil)
	err := tpl.Execute(w, Entries)

	helper.InternalServerErrorCheck(err, w)
}

func entry(c web.C, w http.ResponseWriter, r *http.Request) {

	entry, entryNotFound := model.FindEntry(c.URLParams["id"])

	if entryNotFound {
		NotFound(w, r)
		return
	}

	p := bluemonday.UGCPolicy()
	htmlContent := p.Sanitize(string(blackfriday.MarkdownCommon([]byte(entry.Content))))

	tpl, _ := ace.Load("views/layouts/layout", "views/view", nil)
	err := tpl.Execute(w, map[string]string{"Title": entry.Title, "HtmlContent": htmlContent})

	helper.InternalServerErrorCheck(err, w)
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", nil)
	err := tpl.Execute(w, nil)

	helper.InternalServerErrorCheck(err, w)
}

func postEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	themeId, _ := strconv.Atoi(r.FormValue("entry[theme_id]"))

	// TODO: validation
	if len(title) <= 0 || len(title) > 140 {
		http.Redirect(w, r, "/new", http.StatusNotModified)
		return
	}
	if len(content) < 5 || len(content) > 1000 {
		http.Redirect(w, r, "/new", http.StatusNotModified)
		return
	}

	entry := model.Entry{
		Title:   title,
		Content: content,
		ThemeId: themeId,
	}

	_, err := govalidator.ValidateStruct(entry)
	if err != nil {
		println("error: " + err.Error())
	}

	db.Dbmap.NewRecord(entry)
	db.Dbmap.Create(&entry)

	url := fmt.Sprintf("/%d", entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", nil)
	tpl.Execute(w, nil)
}
