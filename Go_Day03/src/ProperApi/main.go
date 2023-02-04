package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Place struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"location"`
}

type Info struct {
	Name     string
	Total    int
	Places   []Place
	PrevPage int
	NextPage int
	LastPage int
}

type Error struct {
	Page  int
	Value string
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

func FormBody(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
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
			&infoPlaces[i].Address, &infoPlaces[i].Location.Longitude, &infoPlaces[i].Location.Latitude)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	if page > 1364 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error{
			Page:  page,
			Value: "foo",
		})
		return
	}
	json.NewEncoder(w).Encode(Info{
		Name:     "Places",
		Total:    TotalPlaces(),
		Places:   infoPlaces,
		PrevPage: page - 1,
		NextPage: page + 1,
		LastPage: 1364,
	})
}

func main() {
	router := httprouter.New()
	router.GET("/", FormBody)
	log.Fatal(http.ListenAndServe(":8080", router))
}
