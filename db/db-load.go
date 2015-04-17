package db

import (
	"log"
	"os"
)

func DbLoad() {
	env := "development"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}

	if env == "production" {
		DbConnect("production")
	} else {
		DbConnect("development")
	}

	log.Println(env)
}
