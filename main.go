package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	ResponseCode int `json:"response_code"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handles the un-registered routes
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<p>Server up and running!</p>"))
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Message: "Server is up and running!",
			ResponseCode: 200,
		}
		json.NewEncoder(w).Encode(resp)
	})

	serverAddr := ":8080"
	log.Printf("Server running at http://localhost%s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
