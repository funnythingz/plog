package main

import (
	"./db"
	"github.com/shaoshing/train"
	"github.com/zenazn/goji"
	"net/http"
	"regexp"
)

var (
	topHandler       = TopHandler{}
	entryHandler     = EntryHandler{}
	exceptionHandler = ExceptionHandler{}
)

func main() {
	// Database
	db.Connect()

	// Assets
	train.ConfigureHttpHandler(nil)
	goji.Get("/assets/*", http.FileServer(http.Dir("./public/")))

	// Index
	goji.Get("/", topHandler.Index)

	// Entry
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entryHandler.Entry)
	goji.Get("/new", entryHandler.New)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", entryHandler.Create)
	goji.Post(regexp.MustCompile(`^/(?P<id>\d+)/comment$`), entryHandler.AddComment)

	// Exception
	goji.NotFound(exceptionHandler.NotFound)

	// Serve
	goji.Serve()
}
