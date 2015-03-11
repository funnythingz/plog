package main

import (
	"./db"
	"./models"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
	"unicode/utf8"
)

func addComment(c web.C, w http.ResponseWriter, r *http.Request) {
	content := Sanitize(r.FormValue("comment[content]"))

	pp.Println(content)

	entryId, _ := strconv.Atoi(c.URLParams["id"])

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
		errors = append(errors, "input Title must be blank.")
	}
	if utf8.RuneCountInString(content) > 120 {
		errors = append(errors, "input Title maximum is 120 character.")
	}

	if len(errors) > 0 {
		// TODO: error
		pp.Println(errors)
		return
	}

	db.Dbmap.NewRecord(comment)
	db.Dbmap.Create(&comment)
	pp.Println("Create: ", comment)

	url := fmt.Sprintf("/%d", entryId)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
