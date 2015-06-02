package viewmodels

import (
	"../config"
	"../models"
	"fmt"
	"github.com/funnythingz/sunnyday"
	"github.com/garyburd/redigo/redis"
	"github.com/k0kubun/pp"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"log"
	"regexp"
	"strconv"
	"time"
)

type EntryViewModel struct {
	Entry       model.Entry
	Date        string
	HtmlContent string
	Pv          string
	MetaOg      MetaOg
	Flash       []string
}

func (vm *EntryViewModel) Store(entry model.Entry) EntryViewModel {
	p := bluemonday.UGCPolicy()
	htmlContent := p.Sanitize(string(blackfriday.MarkdownCommon([]byte(entry.Content))))

	reg := regexp.MustCompile(`([\s]{2,}|\n)`)

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	entryCreatedAtJST := entry.CreatedAt.In(jst)

	return EntryViewModel{
		Entry:       entry,
		Date:        entryCreatedAtJST.Format(time.ANSIC),
		HtmlContent: htmlContent,
		Pv:          pv(fmt.Sprintf("%d", entry.Id)),
		MetaOg: MetaOg{
			Title: entry.Title,
			Type:  "article",
			//TODO: Url: entry.Id,
			//TODO: Image:  "",
			Description: sunnyday.Truncate(reg.ReplaceAllString(entry.Content, " "), 99),
		},
	}
}

func pv(id string) string {
	conn, err := redis.Dial("tcp", config.Param.Redis.Address)
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("AUTH", config.Param.Redis.AuthPassword)
	if err != nil {
		log.Println("Redis AUTH error")
	}
	defer conn.Close()

	key := fmt.Sprintf("entry_%s", id)

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
