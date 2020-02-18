package chitchat

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// /static/で始まるすべてのリクエストURLに対して、
	// staticを取り去り、ディレクトリpublicを起点として、残った文字列を名前にもファイルを探す
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server {
		Addr:              "0.0.0.0:8080",
		Handler:           mux,
	}
	server.ListenAndServe()
}