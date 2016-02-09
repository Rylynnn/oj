package controller

import (
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

//HandleSubmitCode : Handle user submited code from problem page
func HandleSubmitCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()

	}
}
