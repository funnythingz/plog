package main

import (
	"./db"
	"github.com/shaoshing/train"
	"github.com/zenazn/goji"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"regexp"
)

func main() {

	errcd := Daemon(0, 0)
	if errcd != 0 {
		os.Exit(1)
	}

	os.Chdir("./")

	db.DbDevelopmentConnect()
	train.ConfigureHttpHandler(nil)

	goji.Get("/", top)
	goji.Get("/assets/*", http.FileServer(http.Dir("./public/")))
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entry)
	goji.Get("/new", newEntry)
	goji.Get("/entry", http.RedirectHandler("/", 301))
	goji.Post("/entry", createEntry)

	goji.NotFound(NotFound)

	listener, _ := net.Listen("tcp", "127.0.0.1:8000")
	fcgi.Serve(listener, goji.DefaultMux)
}
