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
