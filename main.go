package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/majorchork/blogapp/handlers"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates *template.Template

type Blogpost struct {
	PostTitle   string
	PostData    string
	WritersName string
	Time        time.Time
	Edit        string
	Delete      string
}

var Blog []Blogpost

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	handlers.Run(router)
	//http.Handle("/", router)
	fmt.Println("Server started....")
	err := http.ListenAndServe(":1769", router)
	if err != nil {
		log.Println(err)
		return
	}

}

func RenderHandler(w http.ResponseWriter, router *http.Request) {
	file, err := templates.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = file.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		return
	}

}
func Test(w http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}

	Title := request.PostForm.Get("post-title")
	Story := request.PostForm.Get("post-data")
	Writer := request.PostForm.Get("writers-name")

	newPost := Blogpost{
		PostTitle:   Title,
		PostData:    Story,
		WritersName: Writer,
		Time:        time.Now(),
	}

	Blog = append(Blog, newPost)
	fmt.Println(Blog)

	temp := template.Must(template.ParseFiles("templates/feed.html"))
	er := temp.Execute(w, newPost)
	if er != nil {
		log.Fatal(er)
	}

	http.Redirect(w, request, "/feed", http.StatusMovedPermanently)
}
