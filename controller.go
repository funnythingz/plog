package main

import (
	"./helper"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
)

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/top", nil)
	err := tpl.Execute(w, map[string]string{"Title": "Welcome"})

	helper.InternalServerErrorCheck(err, w)
}

func article(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/show", nil)
	err := tpl.Execute(w, map[string]string{"Title": "Yeah, " + c.URLParams["basename"] + "!"})

	helper.InternalServerErrorCheck(err, w)
}
