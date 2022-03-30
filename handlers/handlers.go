package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

type news struct {
	Id        string
	PostTitle string
	Story     string
	Writer    string
	Time      string
}

var i news
var j []news
var templates *template.Template

func Run(r *chi.Mux) {
	r.Get("/feed", HomePageHandler)
	r.Get("/form", FormPageHandler)
	r.Post("/feed", PostHomeageHandler)
	//r.Patch("/feed", EditHandler)
	//r.Delete("/feed", DeleteHandler)

}

func HomePageHandler(writer http.ResponseWriter, request *http.Request) {
	file, err := templates.ParseFiles("templates/feed.html")
	if err != nil {
		return
	}
	err = file.Execute(writer, j)
	if err != nil {
		return
	}
}

func FormPageHandler(writer http.ResponseWriter, request *http.Request) {
	file, err := templates.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("error parsing files on index page")
		return
	}
	err = file.Execute(writer, nil)
	if err != nil {
		log.Fatalf("error executing template on index page")
		return
	}
}

func PostHomeageHandler(writer http.ResponseWriter, request *http.Request) {
	Title := request.FormValue("post-title")
	Story := request.FormValue("post-data")
	Writer := request.FormValue("writers-name")
	Time := time.Now().String()
	Id := uuid.NewString()
	i.PostTitle = Title
	i.Story = Story
	i.Writer = Writer
	i.Id = Id
	i.Time = Time
	j = append(j, i)
	fmt.Println(j)

	temp := template.Must(template.ParseFiles("templates/feed.html"))
	temp.Execute(writer, j)

}

func EditHandler(writer http.ResponseWriter, request *http.Request) {

}

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {

}
