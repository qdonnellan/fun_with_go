package hello

import (
  "net/http"
  "html/template"

  "appengine"
)

func init() {
  http.HandleFunc("/", viewHandler)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}