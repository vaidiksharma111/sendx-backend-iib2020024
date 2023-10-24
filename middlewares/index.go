package middlewares

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home.html")
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./configuration.html")
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./crawl.html")
}

func UserLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./user-login.html")
}

func AdminLoginForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./admin-login.html")
}
