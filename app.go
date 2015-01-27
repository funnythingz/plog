package main

import (
	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/", top)
	goji.Get("/:id", hello)
	goji.Serve()
}

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("views/layouts/layout", "views/top", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"Title": "Welcome"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("views/layouts/layout", "views/show", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"Title": "Yeah, " + c.URLParams["id"] + "!"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
