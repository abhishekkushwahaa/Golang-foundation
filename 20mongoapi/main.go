package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhishekkushwahaa/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API");
	fmt.Println("Server is running on port 4000...")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening is running on port 4000...")
}