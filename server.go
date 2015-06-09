package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl map[string]*template.Template

// init is run before the application starts serving.
func init() {
	if tmpl == nil {
		tmpl = make(map[string]*template.Template)
	}
	//baseTmpl := template.Must(template.ParseFiles("templates/base.html"))
	tmpl["home"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	tmpl["register"] = template.Must(template.ParseFiles("templates/register.html", "templates/base.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	hm := Page{"Bienvenue", "contenu de la page qui provient sans doute d'une BDD"}
	err := tmpl["home"].ExecuteTemplate(w, "base", hm)
	if err != nil {
		log.Fatal(err)
	}

	r.ParseForm()                      // Traitement du contenu du formulaire
	fmt.Fprintf(w, "<br />%v", r.Form) // <- On démarre l'insertion dans Neo4j à partir de là
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	hm := Page{"Inscrivez vous", "Hello wrrrrld..."}
	err := tmpl["register"].ExecuteTemplate(w, "base", hm)
	if err != nil {
		log.Fatal(err)
	}
}
