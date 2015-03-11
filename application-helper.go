package main

import (
	"github.com/microcosm-cc/bluemonday"
)

func Sanitize(s string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(string(s))
}
