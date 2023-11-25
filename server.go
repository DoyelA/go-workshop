package main

import (
	"fmt"
	"net/http"
)

func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to GoFr's event!\n"))
}

func main() {
	http.HandleFunc("/event", eventHandler)
	fmt.Printf("Server is running on http://localhost:8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server encountered an error : %s\n", err)
	}

}
