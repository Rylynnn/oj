package controller

import "github.com/unrolled/render"

var (
	//Render :render data var
	Render *render.Render
)

// Init :init controller methods
func init() {
	Render = render.New(render.Options{
		Directory: "templates",
		Layout:    "layout",
	})
}
