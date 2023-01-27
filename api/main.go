package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()
	config.LoadEnv()
	fmt.Println("API is up and running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
