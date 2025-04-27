package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func init() {
	config.LoadEnvs()
}

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
