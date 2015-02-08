package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	_ "github.com/goji/param"
	"github.com/k0kubun/pp"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"strconv"
)

type Paginate struct {
	IsEndpoint  bool
	CurrentPage int
	PrevPage    int
	NextPage    int
}

type TopViewModel struct {
	Entries  []model.Entry
	NotFound bool
	Paginate Paginate
}

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	permit := 5

	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	entries, entriesNotFound, endpoint := model.FindEntriesIndex(permit, page)

	if entriesNotFound {
		NotFound(w, r)
		return
	}

	paginate := Paginate{
		IsEndpoint:  endpoint,
		CurrentPage: page,
	}

	if page > 1 {
		paginate.PrevPage = page - 1
	}

	if endpoint == false {
		paginate.NextPage = page + 1
	}

	TopViewModel := TopViewModel{
		Entries:  entries,
		NotFound: entriesNotFound,
		Paginate: paginate,
	}

	//pp.Println(TopViewModel)

	tpl, _ := ace.Load("views/layouts/layout", "views/top", nil)
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

	tpl, _ := ace.Load("views/layouts/layout", "views/view", nil)
	err := tpl.Execute(w, map[string]string{"Title": entry.Title, "HtmlContent": htmlContent})

	helper.InternalServerErrorCheck(err, w)
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", nil)
	err := tpl.Execute(w, nil)

	helper.InternalServerErrorCheck(err, w)
}

type FormResultData struct {
	Entry model.Entry
	Error []string
}

func createEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	themeId, _ := strconv.Atoi(r.FormValue("entry[theme_id]"))

	Entry := model.Entry{
		Title:   title,
		Content: content,
		ThemeId: themeId,
	}

	if _, err := govalidator.ValidateStruct(Entry); err != nil {
		pp.Println(err.Error())
	}

	Error := []string{}

	// Validation
	if len(title) <= 0 {
		Error = append(Error, "input Title must be blank.")
	}
	if len(title) > 140 {
		Error = append(Error, "input Title maximum is 140 character.")
	}
	if len(content) <= 0 {
		Error = append(Error, "input Content must be blank.")
	}
	if len(content) < 5 || len(content) > 1000 {
		Error = append(Error, "input Content minimum is 5 and maximum is 1000 character.")
	}
	if len(Error) > 0 {
		tpl, _ := ace.Load("views/layouts/layout", "views/new", nil)
		tpl.Execute(w, FormResultData{Entry, Error})
		return
	}

	db.Dbmap.NewRecord(Entry)
	db.Dbmap.Create(&Entry)

	url := fmt.Sprintf("/%d", Entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", nil)
	tpl.Execute(w, nil)
}
