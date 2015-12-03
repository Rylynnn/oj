package middleware

import (
	"net/http"
	"net/url"
	"strings"
)

//Permission :Use to check permission when user touch '/admin/*' (Yet finish)
func Permission(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	u, _ := url.ParseRequestURI(r.URL.String())
	if !strings.HasPrefix(u.Path, "/admin") {
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", 401)
	}
}
