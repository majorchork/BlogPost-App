package handlers

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/majorchork/blogapp/database"
	"html/template"
	"log"
	"net/http"
	"time"
)

//type news struct {
//	Id        string
//	PostTitle string
//	Story     string
//	Writer    string
//	Time      string
//}
//
//var j []news
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
	j := database.Render()

	log.Println("all items in the db: ", j)
	file, err := templates.ParseFiles("templates/feed.html")
	if err != nil {
		log.Fatalf("error parsing files")
	}

	err = file.ExecuteTemplate(writer, "feed.html", j)
	if err != nil {
		log.Fatalf("error %v", err)
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
	var i database.News
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

	database.Create(Id, Writer, Story, Time, Title)
	http.Redirect(writer, request, "/feed", http.StatusMovedPermanently)

}

func EditHandler(writer http.ResponseWriter, request *http.Request) {
	Id := chi.URLParam(request, "Id")
	data := database.EditValue(Id)
	temp := template.Must(template.ParseFiles("templates/update.html"))
	temp.Execute(writer, data)
}

func UpdateHandler(writer http.ResponseWriter, request *http.Request) {
	Id := chi.URLParam(request, "Id")
	data := database.News{
		PostTitle: request.FormValue("post-title"),
		Story:     request.FormValue("post-data"),
		Writer:    request.FormValue("writers-name"),
		Time:      time.Now().Format("Monday, 02-Jan-06"),
		Id:        Id,
	}

	database.Edit(data.Id, data.Writer, data.Story, data.Time, data.PostTitle)

	http.Redirect(writer, request, "/feed", http.StatusMovedPermanently)
}

func ReadHandler(writer http.ResponseWriter, request *http.Request) {
	ID := chi.URLParam(request, "Id")
	p := database.EditValue(ID)

	temp := template.Must(template.ParseFiles("templates/data.html"))
	temp.Execute(writer, p)
}

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	ID := chi.URLParam(request, "Id")
	database.Delete(ID)
	http.Redirect(writer, request, "/feed", http.StatusMovedPermanently)
}
