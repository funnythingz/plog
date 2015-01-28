package main

import (
	"github.com/zenazn/goji"
)

func main() {
	DbConnect()

	goji.Get("/", top)
	goji.Get("/:id", hello)
	goji.Serve()
}
