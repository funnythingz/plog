package main

import (
	"./db"
	"github.com/zenazn/goji"
	"regexp"
)

func main() {
	db.DbDevelopmentConnect()

	goji.Get("/", top)
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), article)

	goji.NotFound(NotFound)
	goji.Serve()
}
