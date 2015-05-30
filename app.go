package main

import (
	"./db"
	"github.com/shaoshing/train"
	"github.com/zenazn/goji"
	"net/http"
	"regexp"
)

var (
	topController       = TopController{}
	entryController     = EntryController{}
	exceptionController = ExceptionController{}
)

func main() {
	// Database
	db.Connect()

	// Assets
	train.ConfigureHttpHandler(nil)
	goji.Get("/assets/*", http.FileServer(http.Dir("./public/")))

	// Index
	goji.Get("/", topController.Index)

	// Entry
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entryController.Entry)
	goji.Get("/new", entryController.New)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", entryController.Create)
	goji.Post(regexp.MustCompile(`^/(?P<id>\d+)/comment$`), entryController.AddComment)

	// Exception
	goji.NotFound(exceptionController.NotFound)

	// Serve
	goji.Serve()
}
