package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// /static/で始まるすべてのリクエストURLに対して、
	// staticを取り去り、ディレクトリpublicを起点として、残った文字列を名前にもファイルを探す
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server {
		Addr:              "0.0.0.0:8080",
		Handler:           mux,
	}
	server.ListenAndServe()
}