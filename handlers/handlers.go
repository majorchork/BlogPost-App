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
	Edit      string
	Delete    string
}

var i news
var j []news
var templates *template.Template

func Run(r *chi.Mux) {
	r.Get("/feed", HomePageHandler)
	r.Get("/form", FormPageHandler)
	r.Post("/feed", PostHomeageHandler)
	//r.Get("/feed/{{.Id}}", EditHandler)
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
	Edit := "EDIT"
	Delete := "DELETE"
	i.PostTitle = Title
	i.Story = Story
	i.Writer = Writer
	i.Id = Id
	i.Time = Time
	i.Edit = Edit
	i.Delete = Delete
	j = append(j, i)
	fmt.Println(j)

	temp := template.Must(template.ParseFiles("templates/feed.html"))
	temp.Execute(writer, j)

}

func EditHandler(writer http.ResponseWriter, request *http.Request) {
	//iterate through struct
	//find maytching i.d
	// j= append(j[0:i], j[i+1:]...)
}

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {

}
