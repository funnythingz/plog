package main

import (
	"./helper"
	"./models"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"strconv"
)

type TopController struct{}

func (_ *TopController) Index(c web.C, w http.ResponseWriter, r *http.Request) {
	permit := 60

	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	entries, nextEntries := model.FindEntriesIndex(permit, page)

	if len(entries) == 0 && page > 1 {
		exceptionController.NotFound(w, r)
		return
	}

	var isFirstpoint bool
	if page == 1 {
		isFirstpoint = true
	}

	var isEndpoint bool
	if len(nextEntries) == 0 {
		isEndpoint = true
	}

	meta := MetaOg{
		Title: "",
		Type:  "website",
		//TODO: Url: "",
		//TODO: Image:  "",
		Description: "plog is a simple diary for people all over the world.",
	}

	TopViewModel := TopViewModel{
		Entries: entries,
		Paginate: Paginate{
			IsFirstpoint: isFirstpoint,
			IsEndpoint:   isEndpoint,
			CurrentPage:  page,
			PrevPage:     page - 1,
			NextPage:     page + 1,
		},
		MetaOg: meta,
	}

	tpl, _ := ace.Load("views/layouts/layout", "views/top", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
	if err := tpl.Execute(w, TopViewModel); err != nil {
		helper.InternalServerErrorCheck(err, w)
	}
}
