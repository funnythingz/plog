package main

import (
	"./db"
	"github.com/shaoshing/train"
	"github.com/zenazn/goji"
	"net/http"
	"regexp"
)

func main() {
	db.DbConnect("development")
	train.ConfigureHttpHandler(nil)

	// Index
	topController := &TopController{}
	goji.Get("/", topController.Index)

	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entry)
	goji.Get("/new", newEntry)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", createEntry)
	goji.Post(regexp.MustCompile(`^/(?P<id>\d+)/comment$`), addComment)

	goji.NotFound(NotFound)
	goji.Serve()
}
