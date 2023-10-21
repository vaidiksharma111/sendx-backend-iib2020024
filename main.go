package main

import (
	"fmt"
	"net/http"
	"project-sendx/handlers"
)

const (
	PORT = 3000
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/user-login", handlers.UserLoginHandler)
	http.HandleFunc("/user-login-form", handlers.UserLoginForm)
	http.HandleFunc("/admin-login-form", handlers.AdminLoginForm)
	http.HandleFunc("/admin-login", handlers.AdminLoginHandler)
	http.HandleFunc("/crawl", handlers.CrawlHandler)
	http.HandleFunc("/config/numWorkers", handlers.UpdateNumWorkers)
	http.HandleFunc("/config/maxCrawlsPerHour", handlers.UpdateMaxCrawlsPerHour)
	http.HandleFunc("/config", handlers.GetConfig)
	http.HandleFunc("/get-config", handlers.GetConfigJSON)

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))

	fmt.Println("Server Running on ", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
