package main

import (
	"./db"
	"github.com/zenazn/goji"
)

func main() {
	dbmap.DbConnect()

	goji.Get("/", top)
	goji.Get("/:id", hello)
	goji.Serve()
}
