package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SohomSaha045/POS_backend/router"
)

func main() {
	fmt.Println("Welcome to POS Backend")
	r := router.Router()
	fmt.Println("Server is getting Started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening to PORT 8080...")

}
