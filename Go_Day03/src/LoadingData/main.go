package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "dbname=golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	query := "CREATE TABLE IF NOT EXISTS Places (" +
		"id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY," +
		"name varchar(255) NOT NULL," +
		"phone varchar(255) NOT NULL," +
		"address varchar(255) NOT NULL," +
		"longitude real NOT NULL, latitude real NOT NULL" +
		")"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}

	// inserting method
	insertQuery := "INSERT INTO Places (id, name, phone, address, longitude, latitude) " +
		"values ($1, $2, $3, $4, $5, $6)"

	data, err := os.Open("/Users/......")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), "\t")
		id, _ := strconv.Atoi(info[0])
		name := info[1]
		address := info[2]
		phone := info[3]
		longitude, _ := strconv.ParseFloat(info[4], 64)
		latitude, _ := strconv.ParseFloat(info[5], 64)
		_, err = db.Exec(insertQuery, id, name, phone, address, longitude, latitude)
	}
}
