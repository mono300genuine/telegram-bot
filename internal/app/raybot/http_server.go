package raybot

import (
	"net/http"
)

func NewHttpServer() *HttpServer {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	server := &http.Server{
		Handler: mux,
	}
	return &HttpServer{server}
}

type HttpServer struct {
	*http.Server
}
