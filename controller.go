package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	_ "github.com/goji/param"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	_ "log"
	"net/http"
)

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/top", nil)
	err := tpl.Execute(w, map[string]string{"Title": "Welcome"})

	helper.InternalServerErrorCheck(err, w)
}

func entry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/view", nil)
	err := tpl.Execute(w, map[string]string{"Title": "Yeah, " + c.URLParams["id"] + "!"})

	helper.InternalServerErrorCheck(err, w)
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", nil)
	err := tpl.Execute(w, nil)

	helper.InternalServerErrorCheck(err, w)
}

func postEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	entry := model.Entry{
		Title:   r.FormValue("entry[title]"),
		Content: r.FormValue("entry[content]"),
		Theme:   1,
	}

	db.Dbmap.NewRecord(entry)
	db.Dbmap.Create(&entry)

	// TODO: validation
	//if err != nil || len(greet.Message) > 140 {
	//    http.Error(w, err.Error(), http.StatusBadRequest)
	//    return
	//}

	url := fmt.Sprintf("/%d", entry.Id)
	http.Redirect(w, r, url, http.StatusCreated)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", nil)
	tpl.Execute(w, nil)
}
