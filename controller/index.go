package controller

import (
	"net/http"
	"strconv"

	"github.com/Miloas/oj/model"
)

const problemsPageNum int = 1

type problemsPageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Pagination   []int
	Problems     []model.Problem
}

//HandleHome :handle "/"
func HandleHome(w http.ResponseWriter, r *http.Request) {
	//假设page是整数,回头改这
	p := 0
	if tmp := r.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("problems")
	count, err := c.Count()
	totalPage := (count + problemsPageNum - 1) / problemsPageNum
	problems := []model.Problem{}
	err = c.Find(nil).Limit(problemsPageNum).Skip(problemsPageNum * p).All(&problems)
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
	result := problemsPageStruct{p, p + 1, p - 1, canNext, canPrevious, pagination, problems}
	Render.HTML(w, http.StatusOK, "index", result)
}
