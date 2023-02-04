package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB
var err error

type Place struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Info struct {
	Total  int
	Places []Place
	Page   int
}

func init() {
	connStr := "postgres://postgres:password@localhost/golang?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("You are now connected to the database.")
}

func TotalPlaces() int {
	infoDB, err := db.Query("SELECT count(*) FROM places")
	if err != nil {
		log.Fatalln(err)
	}
	var total int
	for infoDB.Next() {
		err = infoDB.Scan(&total)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return total
}

func BodyForm(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil || page < 0 {
		w.WriteHeader(400)
		w.Write([]byte("error"))
		return
	}

	rows, err := db.Query("select * from places limit 10 offset $1", page*10)
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	infoPlaces := make([]Place, 10)
	i := 0

	for rows.Next() {

		err = rows.Scan(&infoPlaces[i].Id, &infoPlaces[i].Name, &infoPlaces[i].Phone,
			&infoPlaces[i].Address, &infoPlaces[i].Longitude, &infoPlaces[i].Latitude)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	var info Info

	info.Places = infoPlaces
	info.Total = TotalPlaces()
	info.Page = page

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
		"div": func(a, b int) int {
			return a / b
		},
	}
	t, err := template.New("example.html").Funcs(funcMap).ParseFiles("example.html")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = t.Execute(w, info)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	http.HandleFunc("/", BodyForm)
	http.ListenAndServe(":8080", nil)
}
