package controller

import (
	"net/http"

	"github.com/Miloas/oj/model"

	"gopkg.in/mgo.v2/bson"
)

//HandleProblem : handle "/problem/:id"
func HandleProblem(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("problems")
	var result model.Problem
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		panic(err)
	}
	Render.HTML(w, http.StatusOK, "problem", result)
}
