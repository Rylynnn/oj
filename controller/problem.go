package controller

import (
	"log"
	"net/http"
	"oj/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//HandleProblem : handle "/problem/:id"
func HandleProblem(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	S, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer S.Close()
	c := S.DB("oj").C("problems")
	var result model.Problem
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		panic(err)
	}
	Render.HTML(w, http.StatusOK, "problem", result)
}
