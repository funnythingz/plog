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

	// Entry
	entryController := &EntryController{}
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entryController.Entry)
	goji.Get("/new", entryController.New)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", entryController.Create)
	goji.Post(regexp.MustCompile(`^/(?P<id>\d+)/comment$`), addComment)

	goji.NotFound(NotFound)
	goji.Serve()
}
