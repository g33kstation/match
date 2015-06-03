package main

import (
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

type Page struct {
	Title   string
	Content string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	hm := Page{"Bienvenue", "Hello wrrrrld..."}
	err := tmpl["home"].ExecuteTemplate(w, "base", hm)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(w, "%#v", tmpl["home"])
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
}
