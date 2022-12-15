package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()

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
			tmpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, "Problème de template html", http.StatusInternalServerError)
			}
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

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Response</title>
	<style>
		:root {
			font-family: "Trebuchet MS", serif;
			font-size: 16px;
		}

		.container {
            width: 80%;
            margin: auto;
		}

		h1 {
			text-align: center;
			text-decoration: underline;
		}

		.response div {
            background-color: #E6E6E6FF;
			padding: 10px;
            margin: 20px 0;
		}

		.response h3 {
			text-decoration: underline;
		}

		button {
            width: 100%;
			border: none;
			padding: 20px;
			font-size: 20px;
            background-color: black;
			color: white;
			transition: 300ms ease-in-out background-color;
		}

		button:hover {
            background-color: #5b5b5b;
		}
	</style>
</head>
<body>
<div class="container">
	<h1>Response</h1>

	<div class="response">
		<div>
			<h3>Prénom : </h3>
			<p>%s</p>
		</div>
		<div>
			<h3>Nom : </h3>
			<p>%s</p>
		</div>
		<div>
			<h3>Raison : </h3>
			<p>%s</p>
		</div>
		<div>
			<h3>Message : </h3>
			<p>%s</p>
		</div>
	</div>

	<button onclick="history.back();">Retour</button>
</div>
</body>
</html>
`
