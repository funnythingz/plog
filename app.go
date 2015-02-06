package main

import (
	"github.com/zenazn/goji"
	"regexp"
)

func main() {
	DbDevelopmentConnect()

	goji.Get("/", top)
	goji.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), article)

	goji.NotFound(NotFound)
	goji.Serve()
}
