package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	ResponseCode int `json:"response_code"`
}

func main() {
	http.HandleFunc("/", func(reponseWriter http.ResponseWriter, request *http.Request) {
		// Every un-registered will result in 404
		if request.URL.Path != "/" {
			http.NotFound(reponseWriter, request)
			return
		}
	
		reponseWriter.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(reponseWriter, "<p>Server up and running!</p>")
	})

	http.HandleFunc("/api", func(reponseWriter http.ResponseWriter, request *http.Request) {
		reponseWriter.Header().Set("Content-Type", "application/json")
		resp := Response{
			Message: "Server is up and running!",
			ResponseCode: 200,
		}
		json.NewEncoder(reponseWriter).Encode(resp)
	})

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}