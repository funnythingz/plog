package main

import (
	"./db"
	"github.com/zenazn/goji"
	"net/http"
	"regexp"
)

func main() {
	db.DbDevelopmentConnect()

	goji.Get("/", top)
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entry)
	goji.Get("/new", newEntry)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", postEntry)

	goji.NotFound(NotFound)
	goji.Serve()
}
