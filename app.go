package main

import (
	"github.com/zenazn/goji"
)

func main() {
	goji.Get("/", top)
	goji.Get("/:id", hello)
	goji.Serve()
}
