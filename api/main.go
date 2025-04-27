package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
