package main

import (
	"fmt"
	"net/http"

	"github.com/allrole-ai/backend-ai/routes"
)

func main() {
	http.HandleFunc("/", routes.URL)
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, nil)
}
