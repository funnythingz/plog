package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	_ "github.com/k0kubun/pp"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type EntryController struct{}

func (_ *EntryController) Entry(c web.C, w http.ResponseWriter, r *http.Request) {

	entry, entryNotFound := model.FindEntry(c.URLParams["id"])

	if entryNotFound {
		exceptionController.NotFound(w, r)
		return
	}

	entryViewModel := &EntryViewModel{}

	tpl, _ := ace.Load("views/layouts/layout", "views/view", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
	if err := tpl.Execute(w, entryViewModel.Store(entry)); err != nil {
		helper.InternalServerErrorCheck(err, w)
	}

}

func (_ *EntryController) New(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
	if err := tpl.Execute(w, NewViewModel{Colors: Colors, Theme: "white"}); err != nil {
		helper.InternalServerErrorCheck(err, w)
	}
}

func (_ *EntryController) Create(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	theme := r.FormValue("entry[theme]")

	Entry := model.Entry{
		Title:   title,
		Content: content,
		Theme:   theme,
	}

	if _, err := govalidator.ValidateStruct(Entry); err != nil {
		log.Println(err.Error())
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
		log.Println(err)
		log.Println(errors)
		return
	}

	db.Dbmap.NewRecord(Entry)
	db.Dbmap.Create(&Entry)

	url := fmt.Sprintf("/%d", Entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (_ *EntryController) AddComment(c web.C, w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`([\s]{2,}|\n|^[\s]+$)`)
	space_reg := regexp.MustCompile(`^[\s]+$`)
	content := space_reg.ReplaceAllString(reg.ReplaceAllString(helper.Sanitizer(r.FormValue("comment[content]")), " "), "")

	entryId, _ := strconv.Atoi(c.URLParams["id"])
	url := fmt.Sprintf("/%d", entryId)

	comment := model.Comment{
		Content: content,
		EntryId: entryId,
	}

	if _, err := govalidator.ValidateStruct(comment); err != nil {
		log.Println(err.Error())
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
		entry, entryNotFound := model.FindEntry(c.URLParams["id"])

		if entryNotFound {
			exceptionController.NotFound(w, r)
			return
		}

		entryViewModel := &EntryViewModel{}
		entryViewModel.Flash = errors

		tpl, _ := ace.Load("views/layouts/layout", "views/view", &ace.Options{DynamicReload: true, FuncMap: ViewHelper})
		if err := tpl.Execute(w, entryViewModel.Store(entry)); err != nil {
			helper.InternalServerErrorCheck(err, w)
		}
		return
	}

	db.Dbmap.NewRecord(comment)
	db.Dbmap.Create(&comment)

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
