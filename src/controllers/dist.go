package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego/logs"
)

type DistHandle struct {
	path string
}

func NewDistHandle(p string) *DistHandle {
	return &DistHandle{p}
}

func (t *DistHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	filePath := t.path + upath
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		filePath = t.path + `/index.html`
		http.ServeFile(w, r, filePath)
	} else if err == nil {
		http.ServeFile(w, r, filePath)
	} else {
		logs.Error(err)
		return
	}
}
