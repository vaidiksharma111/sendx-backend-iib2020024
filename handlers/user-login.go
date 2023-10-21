package handlers

import "net/http"

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/crawl.html")
}

func UserLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/user-login.html")
}

func AdminLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/admin-login.html")
}
