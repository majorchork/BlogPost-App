package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/majorchork/blogapp/db"
	"github.com/majorchork/blogapp/handlers"
	"log"
	"net/http"
)

func main() {
	db.OpenDb()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	handlers.Run(router)
	fmt.Println("Server Running....")
	err := http.ListenAndServe(":1769", router)
	if err != nil {
		log.Println(err)
		return
	}

	defer db.DBClient.Close()
}
