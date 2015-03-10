package main

import (
	"github.com/funnythingz/sunnyday"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shaoshing/train"
	"html/template"
)

var AssetsMap = template.FuncMap{
	"javascript_tag": train.JavascriptTag,
	"stylesheet_tag": train.StylesheetTag,
	"truncate": func(s string, c int) string {
		return sunnyday.Truncate(s, c)
	},
	"sanitize": func(s string) string {
		p := bluemonday.UGCPolicy()
		return p.Sanitize(string(s))
	},
}
