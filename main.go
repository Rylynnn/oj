package main

import (
	"net/http"
	"oj/controller"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HandlHome)

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(mux)
	n.Run(":8080")

}
