package main

import (
	"./helper"
	"./models"
	"github.com/funnythingz/sunnyday"
	"github.com/garyburd/redigo/redis"
	"github.com/k0kubun/pp"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	_ "log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

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
		//TODO: Url: entry.Id,
		//TODO: Image:  "",
		Description: sunnyday.Truncate(reg.ReplaceAllString(entry.Content, " "), 99),
	}

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
	if err := tpl.Execute(w, entryViewModel); err != nil {
		helper.InternalServerErrorCheck(err, w)
	}

}
