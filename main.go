package main

import (
	"fmt"
	"net/http"

	"github.com/letanthang/simple_server/handlers"
)

func main() {
	http.HandleFunc("/api/v1/health", handlers.CheckHealth)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
	fmt.Println("Server listent at 9090")
}
