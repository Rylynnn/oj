package controller

import "net/http"

//HandlHome :handle "/"
func HandlHome(w http.ResponseWriter, req *http.Request) {
	Render.HTML(w, http.StatusOK, "index", nil)
}
