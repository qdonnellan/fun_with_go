package hello

import (
  "net/http"
  "html/template"
)

func init() {
  http.HandleFunc("/", handler)
}

type Page struct {
  FirstName string
  LastName string
  Title string
}

var viewTemplate = template.Must(template.ParseFiles("hello/view.html"))

func handler(w http.ResponseWriter, r *http.Request) {
  newPage := Page{
    FirstName: "Julie",
    LastName: "Wonders",
    Title: "A very good title",
  }
  err := viewTemplate.Execute(w , newPage)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
