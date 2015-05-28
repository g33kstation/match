package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Affichage des profils</h1>")
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Formulaire inscription</h1>")
}
