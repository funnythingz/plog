package helper

import (
	"net/http"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InternalServerErrorCheck(e error, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
