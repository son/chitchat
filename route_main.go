package main

import (
	"html/template"
	"net/http"
	"data"
)

// HTMLを生成してResponseWriterに書き出してる
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}

}
