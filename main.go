package main

import (
	"html/template"
	"net/http"
)





var tpl *template.Template

func main()  {
  tpl , _ = tpl.ParseGlob("./templates/*.html")
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/about", aboutHandler)
  http.HandleFunc("/contact", contactHandler)
  http.HandleFunc("/login", loginHandler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
  tpl.ExecuteTemplate(w, "index.html", nil)
}


func aboutHandler(w http.ResponseWriter,r *http.Request)  {
  tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request)  {
  tpl.ExecuteTemplate(w, "contact.html", nil)
}


func loginHandler(w http.ResponseWriter, r *http.Request)  {
  tpl.ExecuteTemplate(w ,"login.html", nil)
}
