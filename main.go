package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://superfile.dev", http.StatusMovedPermanently)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", redirectHandler)

	fmt.Printf("Starting redirect server on port %s\n", port)
	fmt.Println("All traffic will be redirected to https://superfile.dev")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}