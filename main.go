package main

import (
	"fmt"
	"net/http"

	"github.com/marcosfilipe/auto-demo-go-app/internal/calculator"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		result := calculator.Add(2, 3)
		fmt.Fprintf(w, "2 + 3 = %d", result)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
