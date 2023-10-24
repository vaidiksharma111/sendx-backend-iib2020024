package middlewares

import "net/http"

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./configuration.html")
}
