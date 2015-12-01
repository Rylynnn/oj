package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

//Render :render data var
var Render *render.Render

// Init :init controller methods
func Init() {
	Render = render.New(render.Options{
		Directory: "templates",
		Layout:    "layout",
	})
}

//HandlHome :handle "/"
func HandlHome(w http.ResponseWriter, req *http.Request) {
	Render.HTML(w, http.StatusOK, "index", nil)
}
