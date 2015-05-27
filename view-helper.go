package main

import (
	"./helper"
	"github.com/funnythingz/sunnyday"
	"github.com/shaoshing/train"
	"html/template"
)

var ViewHelper = template.FuncMap{
	"javascript_tag": train.JavascriptTag,
	"stylesheet_tag": train.StylesheetTag,
	"truncate": func(s string, c int) string {
		return sunnyday.Truncate(s, c)
	},
	"sanitize": func(s string) string {
		return helper.Sanitizer(s)
	},
}
