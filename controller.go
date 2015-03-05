package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/funnythingz/sunnyday"
	"github.com/garyburd/redigo/redis"
	"github.com/k0kubun/pp"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/shaoshing/train"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"html/template"
	_ "log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"
)

type Paginate struct {
	IsEndpoint   bool
	IsFirstpoint bool
	CurrentPage  int
	PrevPage     int
	NextPage     int
}

type MetaOg struct {
	Title       string
	Type        string
	Url         string
	Image       string
	Description string
}

type TopViewModel struct {
	Entries  []model.Entry
	Paginate Paginate
	Theme    string
	MetaOg   MetaOg
}

var AssetsMap = template.FuncMap{
	"javascript_tag": train.JavascriptTag,
	"stylesheet_tag": train.StylesheetTag,
	"truncate": func(s string, c int) string {
		return sunnyday.Truncate(s, c)
	},
	"url_encode": func(s string) string {
		return url.QueryEscape(s)
	},
}

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	permit := 60

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

	meta := MetaOg{
		Title: "plog is a simple diary for people all over the world.",
		Type:  "website",
		Url:   "http://plog.link",
		//TODO: Image:  "",
		Description: "plog is a simple diary for people all over the world.",
	}

	TopViewModel := TopViewModel{
		Entries:  entries,
		Paginate: paginate,
		Theme:    "",
		MetaOg:   meta,
	}

	pp.Println(paginate)

	tpl, _ := ace.Load("views/layouts/layout", "views/top", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, TopViewModel)

	helper.InternalServerErrorCheck(err, w)
}

type EntryViewModel struct {
	Title       string
	Date        string
	HtmlContent string
	Theme       string
	MetaOg      MetaOg
	PageView    string
}

func pageView(id string) string {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	key := "entry_" + id

	count := 0
	c, err := redis.String(conn.Do("GET", key))
	if err != nil {
		pp.Println("key not found")
	} else {
		count, _ = strconv.Atoi(c)
	}

	count = count + 1
	conn.Do("SET", key, count)
	resultCount, _ := redis.String(conn.Do("GET", key))

	return resultCount
}

func entry(c web.C, w http.ResponseWriter, r *http.Request) {

	entry, entryNotFound := model.FindEntry(c.URLParams["id"])

	if entryNotFound {
		NotFound(w, r)
		return
	}

	p := bluemonday.UGCPolicy()
	htmlContent := p.Sanitize(string(blackfriday.MarkdownCommon([]byte(entry.Content))))

	reg := regexp.MustCompile(`([\s]{2,}|\n)`)
	meta := MetaOg{
		Title: entry.Title,
		Type:  "article",
		Url:   "http://plog.link/" + strconv.Itoa(entry.Id),
		//TODO: Image:  "",
		Description: sunnyday.Truncate(reg.ReplaceAllString(entry.Content, " "), 99),
	}

	pp.Println(entry)
	pp.Println(meta)

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	entryCreatedAtJST := entry.CreatedAt.In(jst)

	entryViewModel := EntryViewModel{
		Title:       entry.Title,
		Date:        entryCreatedAtJST.Format(time.ANSIC),
		HtmlContent: htmlContent,
		Theme:       entry.Theme,
		MetaOg:      meta,
		PageView:    pageView(c.URLParams["id"]),
	}

	tpl, _ := ace.Load("views/layouts/layout", "views/view", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, entryViewModel)

	helper.InternalServerErrorCheck(err, w)
}

var Colors []string = []string{"white", "black", "pink", "blue", "sky", "green", "purple", "yellow", "lime"}

type NewViewModel struct {
	Entry  model.Entry
	Error  []string
	Theme  string
	MetaOg MetaOg
	Colors []string
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	err := tpl.Execute(w, NewViewModel{Colors: Colors, Theme: "white"})

	helper.InternalServerErrorCheck(err, w)
}
func createEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	theme := r.FormValue("entry[theme]")

	pp.Println(title, content, theme)

	Entry := model.Entry{
		Title:   title,
		Content: content,
		Theme:   theme,
	}

	if _, err := govalidator.ValidateStruct(Entry); err != nil {
		pp.Println(err.Error())
	}

	errors := []string{}

	// Validation
	if utf8.RuneCountInString(title) <= 0 {
		errors = append(errors, "input Title must be blank.")
	}
	if utf8.RuneCountInString(title) > 50 {
		errors = append(errors, "input Title maximum is 50 character.")
	}
	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 5 and maximum is 1000 character.")
	}

	newViewModel := NewViewModel{Entry: Entry, Error: errors, Theme: theme, MetaOg: MetaOg{}, Colors: Colors}

	if len(errors) > 0 {
		tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
		err := tpl.Execute(w, newViewModel)
		pp.Println(err)
		pp.Println(errors)
		return
	}

	db.Dbmap.NewRecord(Entry)
	db.Dbmap.Create(&Entry)
	pp.Println("Create: ", Entry)

	url := fmt.Sprintf("/%d", Entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	tpl.Execute(w, nil)
}
