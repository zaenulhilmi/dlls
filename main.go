package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8010"
	fmt.Printf("Starting server on http://localhost:%s\n", port)

	mux := http.NewServeMux()

	corsMux := enaableCORS(mux)


	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsMux)

	if err != nil {
		fmt.Println("Error starting server")
		fmt.Println(err)
	} else {
		fmt.Printf("Server running on port %s\n", port)
	}
}

func enaableCORS(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})
}
