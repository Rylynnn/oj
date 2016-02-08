package controller

import (
	"log"
	"net/http"
	"github.com/Miloas/oj/model"
	"strconv"

	"gopkg.in/mgo.v2"
)

const pageNum int = 1

type pageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Pagination   []int
	Problems     []model.Problem
}

//HandleHome :handle "/"
func HandleHome(w http.ResponseWriter, req *http.Request) {
	//假设page是整数,回头改这
	p := 0
	if tmp := req.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	S, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer S.Close()
	c := S.DB("oj").C("problems")
	count, err := c.Count()
	totalPage := (count + pageNum - 1) / pageNum
	problems := []model.Problem{}
	err = c.Find(nil).Limit(pageNum).Skip(pageNum * p).All(&problems)
	if err != nil {
		panic(err)
	}
	pagination := []int{}
	for i := 0; i < totalPage; i++ {
		pagination = append(pagination, i)
	}
	canNext, canPrevious := false, false
	if p+1 <= totalPage {
		canNext = true
	}
	if p-1 >= 0 {
		canPrevious = true
	}
	result := pageStruct{p, p + 1, p - 1, canNext, canPrevious, pagination, problems}
	Render.HTML(w, http.StatusOK, "index", result)
}
