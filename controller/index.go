package controller

import (
	"net/http"
	"github.com/ryomak/kurisuuu/core"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	device := core.SetDevice(r.UserAgent())
	if device == "pc" {
		tmpl := template.Must(template.ParseFiles("./public/pc_index.html"))
		tmpl.Execute(w, "")
		return
	}
	tmpl := template.Must(template.ParseFiles("./public/sp_index.html"))
	tmpl.Execute(w,"")
}

