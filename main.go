package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/harshRishi/mongoapis/Router"
)

func main() {
	fmt.Println("MongoDb APIs")
	r := router.Router()

	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server is up and running on port: 4000")
}
