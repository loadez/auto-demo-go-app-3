package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example/auto-demo-go-app/todo"
)

func main() {
	store := todo.NewStore()
	handler := todo.NewHandler(store, "templates")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.List)
	mux.HandleFunc("POST /add", handler.Add)
	mux.HandleFunc("POST /toggle/{id}", handler.Toggle)
	mux.HandleFunc("POST /delete/{id}", handler.Delete)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
