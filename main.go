package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	//http.ListenAndServe(":1212", http.FileServer(http.Dir("public")))

	templates := populateTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedfile := r.URL.Path[1:]
		t := templates.Lookup(requestedfile + ".html")

		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})

	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":1212", nil)
}

func populateTemplate() *template.Template {

	result := template.New("templates")
	const basepath = "templates"
	template.Must(result.ParseGlob(basepath + "/*.html"))

	return result

}
