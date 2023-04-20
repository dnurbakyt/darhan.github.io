package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func whenBornPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		age1, err := strconv.Atoi(age)
		if err != nil {
			fmt.Fprintf(w, "Non number!")
			return
		}
		year := 2023 - age1
		output := map[string]interface{}{"name": name, "year": year}
		tmpl, _ := template.ParseFiles("static/form.html")
		tmpl.Execute(w, output)

	} else {
		tmpl, _ := template.ParseFiles("static/form.html")
		tmpl.Execute(w, nil)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/whenborn", whenBornPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
