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
	Read      string
}

var i news
var j []news
var templates *template.Template

func Run(r *chi.Mux) {
	r.Get("/feed", HomePageHandler)
	r.Get("/form", FormPageHandler)
	r.Post("/feed", PostHomeageHandler)
	r.Get("/edit/{Id}", EditHandler)
	r.Post("/update/{Id}", UpdateHandler)
	r.Get("/del/{Id}", DeleteHandler)
	r.Get("/read/{Id}", ReadHandler)

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
	Read := "READ"
	i.PostTitle = Title
	i.Story = Story
	i.Writer = Writer
	i.Id = Id
	i.Time = Time
	i.Edit = Edit
	i.Delete = Delete
	i.Read = Read
	j = append(j, i)
	fmt.Println(j)

	temp := template.Must(template.ParseFiles("templates/feed.html"))
	temp.Execute(writer, j)

}

func EditHandler(writer http.ResponseWriter, request *http.Request) {
	Id := chi.URLParam(request, "Id")
	var data news
	for _, val := range j {
		if val.Id == Id {
			data = val
		}
	}
	temp := template.Must(template.ParseFiles("templates/update.html"))
	temp.Execute(writer, data)
}

func UpdateHandler(writer http.ResponseWriter, request *http.Request) {
	Id := chi.URLParam(request, "Id")
	data := news{
		PostTitle: request.FormValue("post-title"),
		Story:     request.FormValue("post-story"),
		Writer:    request.FormValue("writers-name"),
		Time:      time.Now().String(),
		Id:        uuid.NewString(),
		Edit:      "EDIT",
		Delete:    "DELETE",
		Read:      "READ",
	}
	for i, _ := range j {
		if j[i].Id == Id {
			j[i] = data
			fmt.Println(j[i])
			break
		}
	}
	http.Redirect(writer, request, "/feed?234567yuhgbvcdsw3", http.StatusMovedPermanently)
}
func ReadHandler(writer http.ResponseWriter, request *http.Request) {
	ID := chi.URLParam(request, "Id")
	var p news
	for ind, _ := range j {
		if ID == j[ind].Id {
			p = j[ind]
			break
		} else if ind == len(j)-1 && ID != j[ind].Id {
			log.Fatalf("user Id: %s is invalid", ID)
			return
		}
	}
	temp := template.Must(template.ParseFiles("templates/data.html"))
	temp.Execute(writer, p)
}

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	ID := chi.URLParam(request, "Id")
	for ind, _ := range j {
		if ID == j[ind].Id {
			j = append(j[:ind], j[ind+1:]...)
			break
		} else if ind == len(j)-1 && ID != j[ind].Id {
			log.Fatalf("user Id: %s is invalid", ID)
			return
		}
	}
	http.Redirect(writer, request, "/feed", 302)
}
