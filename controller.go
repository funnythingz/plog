package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	_ "github.com/k0kubun/pp"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/shaoshing/train"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"
)

type Paginate struct {
	IsEndpoint   bool
	IsFirstpoint bool
	CurrentPage  int
	PrevPage     int
	NextPage     int
}

type TopViewModel struct {
	Entries  []model.Entry
	Paginate Paginate
	Theme    string
}

var AssetsMap = template.FuncMap{
	"javascript_tag": train.JavascriptTag,
	"stylesheet_tag": train.StylesheetTag,
}

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	permit := 20

	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	entries, nextEntries := model.FindEntriesIndex(permit, page)

	if len(entries) == 0 && page > 1 {
		NotFound(w, r)
		return
	}

	var isFirstpoint bool
	if page == 1 {
		isFirstpoint = true
	}

	var isEndpoint bool
	if len(nextEntries) == 0 {
		isEndpoint = true
	}

	paginate := Paginate{
		IsFirstpoint: isFirstpoint,
		IsEndpoint:   isEndpoint,
		CurrentPage:  page,
		PrevPage:     page - 1,
		NextPage:     page + 1,
	}

	TopViewModel := TopViewModel{
		Entries:  entries,
		Paginate: paginate,
		Theme:    "",
	}

	log.Println(paginate)

	tpl, _ := ace.Load("views/layouts/layout", "views/top", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, TopViewModel)

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

	tpl, _ := ace.Load("views/layouts/layout", "views/view", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, map[string]string{"Title": entry.Title, "HtmlContent": htmlContent, "Theme": entry.Theme})

	helper.InternalServerErrorCheck(err, w)
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, nil)

	helper.InternalServerErrorCheck(err, w)
}

type FormResultData struct {
	Entry model.Entry
	Error []string
	Theme string
}

func createEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	theme := r.FormValue("entry[theme]")

	Entry := model.Entry{
		Title:   title,
		Content: content,
		Theme:   theme,
	}

	if _, err := govalidator.ValidateStruct(Entry); err != nil {
		log.Println(err.Error())
	}

	Error := []string{}

	// Validation
	if utf8.RuneCountInString(title) <= 0 {
		Error = append(Error, "input Title must be blank.")
	}
	if utf8.RuneCountInString(title) > 50 {
		Error = append(Error, "input Title maximum is 50 character.")
	}
	if utf8.RuneCountInString(content) <= 0 {
		Error = append(Error, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		Error = append(Error, "input Content minimum is 5 and maximum is 1000 character.")
	}
	if len(Error) > 0 {
		tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
		tpl.Execute(w, FormResultData{Entry, Error, "white"})
		return
	}

	db.Dbmap.NewRecord(Entry)
	db.Dbmap.Create(&Entry)
	log.Println("Create: ", Entry)

	url := fmt.Sprintf("/%d", Entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	tpl.Execute(w, nil)
}
