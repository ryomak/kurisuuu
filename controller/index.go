package controller

import (
	"net/http"
	"github.com/ryomak/kurisuuu/core"
)

func Index(w http.ResponseWriter, r *http.Request) {
	device := core.SetDevice(r.UserAgent())
	if device == "pc" {
		http.ServeFile(w, r, "public/pc_index.html")
		return
	}
	http.ServeFile(w, r, "public/sp_index.html")
}