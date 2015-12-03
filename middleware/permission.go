package middleware

import "net/http"

//Permission :Use to check permission when user touch '/admin/*'
func Permission(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//明天在写这坨代码 - -
	if r.URL.Query().Get("password") == "secret123" {
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", 401)
	}
}
