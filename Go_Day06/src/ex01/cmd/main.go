package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"go_day06/ex01/pkg/models/postgresql"
	"log"
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

type application struct {
	article *postgresql.ArticleModel
}

func main() {
	port := flag.String("p", ":8888", "server port")
	dbName := flag.String("db", "postgres", "DB name")
	dbSSL := flag.String("ssl", "disable", "DB SSL mode")
	flag.Parse()

	connStr := fmt.Sprintf("dbname=%s sslmode=%s", *dbName, *dbSSL)
	db, err := openDB(connStr)
	if err != nil {
		log.Fatalf("Err DB! %v", err)
	}
	defer db.Close()

	app := &application{
		article: &postgresql.ArticleModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *port,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal("ListenAndServe: ", err)
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}
	return f, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
