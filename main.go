package main

import (
	"fmt"
	"net/http"
	middlewares "project-sendx/middlewares"
)

const (
	PORT = 3000
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))
	http.HandleFunc("/", middlewares.IndexHandler)
	http.HandleFunc("/admin-login-form", middlewares.AdminLoginForm)
	http.HandleFunc("/admin-login", middlewares.AdminLoginHandler)
	http.HandleFunc("/user-login", middlewares.UserLoginHandler)
	http.HandleFunc("/user-login-form", middlewares.UserLoginForm)
	http.HandleFunc("/config/numWorkers", middlewares.UpdateNumWorkers)
	http.HandleFunc("/get-config", middlewares.GetConfigJSON)
	http.HandleFunc("/config/maxCrawlsPerHour", middlewares.UpdateMaxCrawlsPerHour)
	http.HandleFunc("/crawl", middlewares.CrawlHandler)
	http.HandleFunc("/config", middlewares.GetConfig)

	fmt.Println("Server Running on ", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
