package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", registrationPage)
	http.HandleFunc("/", homePage)

	log.Println("Listening on 0.0.0.0:7700...")
	log.Fatal(http.ListenAndServe(":7700", nil))
}
