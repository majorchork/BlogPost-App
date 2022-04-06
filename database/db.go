package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type News struct {
	Id        string
	PostTitle string
	Story     string
	Writer    string
	Time      string
}

var DBClient *sql.DB

func OpenDb() {
	db, err := sql.Open("mysql", "root:Login1234@tcp(localhost:3306)/BLOG")
	if err != nil {
		log.Fatalf(" error connecting to database: %v", err)
	}

	DBClient = db

}
func Render() []News {
	var J []News
	query := "SELECT * FROM blogpost"
	file, err := DBClient.Query(query)
	if err != nil {
		log.Fatalf("error loading db")
	}
	defer file.Close()
	for file.Next() {
		var i News
		err = file.Scan(&i.Id, &i.Story, &i.PostTitle, &i.Writer, &i.Time)
		if err != nil {
			log.Fatalf("error scaning file %v", err)
		}
		J = append(J, i)
	}
	return J
}
func Create(Id, Writer, Story, Time, Title string) {
	var insert *sql.Stmt
	OpenDb()
	insert, err := DBClient.Prepare("INSERT INTO `BLOG`.`blogpost`(`client_id`, `writers_name`, `story`, `time`, `post_title`) VALUES (?, ?, ?, ?, ?);")
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
func EditValue(Id string) News {
	data := DBClient.QueryRow("SELECT * FROM  `BLOG`.`blogpost` WHERE client_id = ?;", Id)
	var fields News

	err := data.Scan(&fields.Id, &fields.Writer, &fields.Story, &fields.Time, &fields.Time)
	if err != nil {
		log.Fatalf("error at: %v", err)
	}
	return fields
}
func Edit(Writer, Story, Time, PostTitle, Id string) {
	updateQuery := "UPDATE `blogpost` SET `writers_name`=?, `story`=?, `time`=?, `post_title`=? WHERE(`client_id`=?);"

	query, err := DBClient.Prepare(updateQuery)
	if err != nil {
		log.Fatalf("error executing query: %v", err)
	}
	defer query.Close()
	result, err := query.Exec(Writer, Story, Time, PostTitle, Id)
	rowsAffected, _ := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		log.Fatalf("error updating DB :%v", err)
	}
}
func Delete(ID string) {
	log.Println("entering db query")
	deleteQuery, err := DBClient.Prepare("DELETE FROM `BLOG`.`blogpost` WHERE(`client_id`=?);")
	log.Println("execution db query")
	if err != nil {
		log.Fatalf("error executing DB Query")
	}
	defer deleteQuery.Close()

	result, err := deleteQuery.Exec(ID)

	rowsAffected, _ := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		log.Fatalf("error deleting from DB")
	}

	log.Println("done with db query")
}
