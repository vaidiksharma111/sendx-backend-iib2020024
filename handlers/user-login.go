package handlers

import "net/http"

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./crawl.html")
}

func UserLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./user-login.html")
}

func AdminLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./admin-login.html")
}
