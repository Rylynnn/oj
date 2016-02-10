package controller

import "net/http"

//HandleSubmitCode : Handle user submited code from problem page
func HandleSubmitCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c := RedisPool.Get()
		defer c.Close()
		//	c.Do("LPUSH", "judgeQueue", "test")
	}
}
