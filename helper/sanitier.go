package helper

import (
	"github.com/microcosm-cc/bluemonday"
)

func Sanitizer(s string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(string(s))
}
