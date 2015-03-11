package main

import (
	"./db"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/k0kubun/pp"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
	"unicode/utf8"
)

func createEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	theme := r.FormValue("entry[theme]")

	pp.Println(title, content, theme)

	Entry := model.Entry{
		Title:   title,
		Content: content,
		Theme:   theme,
	}

	if _, err := govalidator.ValidateStruct(Entry); err != nil {
		pp.Println(err.Error())
	}

	errors := []string{}

	// Validation
	if utf8.RuneCountInString(title) <= 0 {
		errors = append(errors, "input Title must be blank.")
	}
	if utf8.RuneCountInString(title) > 50 {
		errors = append(errors, "input Title maximum is 50 character.")
	}
	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 5 and maximum is 1000 character.")
	}

	newViewModel := NewViewModel{Entry: Entry, Error: errors, Theme: theme, MetaOg: MetaOg{}, Colors: Colors}

	if len(errors) > 0 {
		tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
		err := tpl.Execute(w, newViewModel)
		pp.Println(err)
		pp.Println(errors)
		return
	}

	db.Dbmap.NewRecord(Entry)
	db.Dbmap.Create(&Entry)
	pp.Println("Create: ", Entry)

	url := fmt.Sprintf("/%d", Entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
