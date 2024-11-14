package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ver := os.Getenv("APP_VERSION")
	if ver == "" {
		ver = "Green"
	}
	fmt.Fprintf(w, "version: %s", ver)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed", err)
	}
}
