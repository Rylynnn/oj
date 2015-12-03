package main

import (
	"net/http"
	"oj/controller"

	"github.com/codegangsta/negroni"
)

func main() {
	controller.Init()
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HandlHome)
	//add static file server for include static files
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(mux)
	n.Run(":8080")
}
