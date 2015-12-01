package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

var Render *render.Render

func Init() {
	Render = render.New()
}

func HandlHome(w http.ResponseWriter, req *http.Request) {
	Render.HTML(w, http.StatusOK, "index", nil)
}
