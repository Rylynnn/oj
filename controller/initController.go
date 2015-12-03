package controller

import (
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
