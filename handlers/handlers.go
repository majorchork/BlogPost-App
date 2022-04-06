package handlers

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/majorchork/blogapp/db"
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
	//info, err := templates.ParseFiles("templates/feed.html")
	//if err != nil {
	//	log.Fatalf("error parsing files")
	//}
	query := "SELECT * FROM BLOG"
	file, err := db.DBClient.Query(query)
	if err != nil {
		log.Fatalf("error loading db")
	}
	defer file.Close()
	for file.Next() {
		var i news
		err = file.Scan(&i.PostTitle, &i.Story, i.PostTitle, i.Writer, i.Time)
		if err != nil {
			log.Fatalf("error scaning file")
		}
		j = append(j, i)
	}
	templates.ExecuteTemplate(writer, "templates/feed.html", j)

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
	var i news
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
	var insert *sql.Stmt
	db.OpenDb()
	insert, err := db.DBClient.Prepare("INSERT INTO `BLOG`.`blogpost`(`client_id`, `writers_name`, `story`, `time`, `post_title`) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatalf("error inserting")
	}
	defer insert.Close()
	result, err := insert.Exec(Id, Writer, Story, Time, Title)

	rowsAffected, _ := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		log.Fatalf("error inserting rows please check all fields")
	}
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
		Story:     request.FormValue("post-data"),
		Writer:    request.FormValue("writers-name"),
		Time:      time.Now().Format("Monday, 02-Jan-06"),
		Id:        uuid.NewString(),
	}
	for i, _ := range j {
		if j[i].Id == Id {
			j[i] = data
			fmt.Println(j[i])
			break
		}
	}
	updateQuery := "UPDATE `BLOG`.`blogpost` SET `writers_name`=?, `story`=?, `time`=?, `post_title`=? WHERE(`client_id`=?);"

	query, err := db.DBClient.Prepare(updateQuery)
	if err != nil {
		log.Fatalf("error executing query")
	}
	defer query.Close()
	result, err := query.Exec(data.Id, data.Writer, data.Story, data.Time, data.PostTitle)
	rowsAffected, _ := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		log.Fatalf("error updating DB")
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
	deleteQuery, err := db.DBClient.Prepare("DELETE FROM `BLOG`.`blogpost` WHERE(`client_id`=?);")
	if err != nil {
		log.Fatalf("error executing DB Query")
		defer deleteQuery.Close()

		result, err := deleteQuery.Exec(ID)

		rowsAffected, _ := result.RowsAffected()
		if err != nil || rowsAffected != 1 {
			log.Fatalf("error deleting from DB")
		}

	}
	http.Redirect(writer, request, "/feed", 302)
}
