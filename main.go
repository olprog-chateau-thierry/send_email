package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"
)

//go:embed templates/*
var html embed.FS

//go:embed static/*
var static embed.FS

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFS(html, "templates/index.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				return
			}
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
			tmpl := template.Must(template.ParseFS(html, "templates/form.html"))

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
	log.Printf("Listening on %s...\n", port)
	log.Fatalln(http.ListenAndServe(port, mux))
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
