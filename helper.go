package main

import (
	"net/http"
)

var (
	helper = &Helper{}
)

type Helper struct{}

func (h *Helper) Check(e error) {
	if e != nil {
		panic(e)
	}
}

func (h *Helper) InternalServerErrorCheck(e error, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
