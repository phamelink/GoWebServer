package server

import (
	"log"
	"net/http"
)

func StartServer() {
	server := http.Server{Addr: ":8000"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homepage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signup_account", signupAccount)
	log.Fatal(server.ListenAndServe())
}
