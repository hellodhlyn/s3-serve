package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", serve)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server listing 0.0.0.0:" + port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
