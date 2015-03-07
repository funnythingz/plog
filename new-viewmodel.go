package main

import (
	"./helper"
	"./models"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
)

type NewViewModel struct {
	Entry  model.Entry
	Error  []string
	Theme  string
	MetaOg MetaOg
	Colors []string
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	if err := tpl.Execute(w, NewViewModel{Colors: Colors, Theme: "white"}); err != nil {
		helper.InternalServerErrorCheck(err, w)
	}
}
