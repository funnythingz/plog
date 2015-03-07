package main

import (
	"github.com/yosssi/ace"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", &ace.Options{DynamicReload: true, FuncMap: AssetsMap})
	tpl.Execute(w, nil)
}
