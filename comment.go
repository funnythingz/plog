package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/k0kubun/pp"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"net/http"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func addComment(c web.C, w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`([\s]{2,}|\n|^[\s]+$)`)
	space_reg := regexp.MustCompile(`^[\s]+$`)
	content := space_reg.ReplaceAllString(reg.ReplaceAllString(Sanitize(r.FormValue("comment[content]")), " "), "")

	pp.Println(content)

	entryId, _ := strconv.Atoi(c.URLParams["id"])
	url := fmt.Sprintf("/%d", entryId)

	comment := model.Comment{
		Content: content,
		EntryId: entryId,
	}

	if _, err := govalidator.ValidateStruct(comment); err != nil {
		pp.Println(err.Error())
	}

	errors := []string{}

	// Validation
	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Comment must be blank.")
	}
	if utf8.RuneCountInString(content) > 120 {
		errors = append(errors, "input Comment maximum is 120 character.")
	}

	if len(errors) > 0 {
		// TODO: error
		entry, entryNotFound := model.FindEntry(c.URLParams["id"])

		if entryNotFound {
			NotFound(w, r)
			return
		}

		entryViewModel := StoreEntryViewModel(entry)
		entryViewModel.Flash = errors
		pp.Println(entryViewModel)

		tpl, _ := ace.Load("views/layouts/layout", "views/view", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
		if err := tpl.Execute(w, entryViewModel); err != nil {
			helper.InternalServerErrorCheck(err, w)
		}
		return
	}

	db.Dbmap.NewRecord(comment)
	db.Dbmap.Create(&comment)
	pp.Println("Create: ", comment)

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
