package main

import (
	"net/http"

	"github.com/Miloas/oj/controller"
	"github.com/Miloas/oj/middleware"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/problem", controller.HandleProblem)
	mux.HandleFunc("/problem/add", controller.HandleAddProblem)
	mux.HandleFunc("/problem/remove", controller.HandleRemoveProblem)
	mux.HandleFunc("/problem/update", controller.HandleUpdateProblem)
	mux.HandleFunc("/problem/submit", controller.HandleSubmitCode)
	//display problems list
	mux.HandleFunc("/", controller.HandleHome)
	//add static file server for include static files
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middleware.Permission),
		negroni.NewLogger(),
	)
	n.UseHandler(mux)
	n.Run(":8080")
}
