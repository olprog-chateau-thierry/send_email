package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"
)

//go:embed html/*
var html embed.FS

//go:embed html/images
var images embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(html)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFS(html, "html/index.html"))
			tmpl.Execute(w, nil)
		} else {
			log.Println("error")
			http.Error(w, "Ceci n'est pas la bonne Méthode HTTP", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/send_email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Problème de formulaire", http.StatusNotImplemented)
			}
			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			reason := r.FormValue("reason")
			message := r.FormValue("message")

			log.Println(firstname, lastname, reason, message)
			w.Header().Add("Content-Type", "text/html")
			tmpl := template.Must(template.ParseFS(html, "html/form.html"))

			tmpl.Execute(w, map[string]string{
				"firstname": firstname,
				"lastname":  lastname,
				"reason":    getReason(reason),
				"message":   message,
			})
		} else {
			log.Println("error")
			http.Error(w, "Ceci n'est pas la bonne Méthode HTTP", http.StatusMethodNotAllowed)
		}
	})

	log.Fatalln(http.ListenAndServe("localhost:8080", mux))
}

func getReason(str string) string {
	switch str {
	case "admin":
		return "Problème administratif"
	case "tech":
		return "Problème Technique"
	case "juri":
		return "Problème Juridique"
	case "other":
		return "Autres"
	}
	return ""
}
