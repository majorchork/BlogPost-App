package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates *template.Template

type blogpost struct {
	PostTitle   string
	PostData    string
	WritersName string
	Time        time.Time
}

var blog blogpost

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := chi.NewRouter()
	router.Get("/", RenderHandler)
	router.Post("/", PostHandler)
	http.Handle("/", router)
	fmt.Println("Server started....")
	http.ListenAndServe(":1759", nil)
}
func RenderHandler(w http.ResponseWriter, router *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", blog)
	if err != nil {
		log.Fatal(err)
	}
}

func PostHandler(w http.ResponseWriter, router *http.Request) {
	Title := router.FormValue("post-title")
	Story := router.FormValue("post-data")
	Writer := router.FormValue("writers-name")
	blog = blogpost{
		PostTitle:   Title,
		PostData:    Story,
		WritersName: Writer,
		Time:        time.Time{},
	}
	err := templates.ExecuteTemplate(w, "index.html", blog)
	if err != nil {
		log.Fatal(err)
	}
}
