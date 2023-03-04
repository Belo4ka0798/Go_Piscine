package main

import (
	"errors"
	"fmt"
	models "go_day06/ex01/pkg"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.showHome)
	mux.HandleFunc("/admin", app.adminLogin)
	mux.HandleFunc("/admin/menu", app.adminMenu)
	mux.HandleFunc("/articles", app.showArticle)
	//mux.HandleFunc("/articles/create", app.createArticle)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("../ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}

func (app *application) showHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s, err := app.article.AllArticles()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	//for _, article := range s {
	//	fmt.Fprintf(w, "%d\n%s\n%s\n%v\n", article.ID, article.Title, article.Content, article.Created)
	//}
	data := &templateData{Articles: s}
	files := []string{
		"../ui/html/home.page.tmpl",
		"../ui/html/base.layout.tmpl",
		"../ui/html/footer.partial.tmpl",
		"../ui/static/css/main.css",
		"../ui/static/js/main.js",
	}
	// Show home.page.tmpl
	ts, err := template.ParseFiles(files...)
	if err != nil {
		w.Write([]byte("Ошибка загрузки страницы!\n"))
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//func (app *application) createArticle(w http.ResponseWriter, r *http.Request) {
//	id, err := app.article.Insert(title, content)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Fprintf(w, "Статья создана!\n Ее ID = %v", id)
//}

func (app *application) showArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	s, err := app.article.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			fmt.Fprintf(w, "404 Not found!")
			return
		}
		log.Println("Error rows")
		return
	}
	data := &templateData{Article: s}
	files := []string{
		"../ui/html/show.page.tmpl",
		"../ui/html/base.layout.tmpl",
		"../ui/html/footer.partial.tmpl",
	}

	// Парсинг файлов шаблонов...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Fprintf(w, "Error!")
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}

	// Отображаем весь вывод на странице.
	//fmt.Fprintf(w, "%d\n%s\n%s\n%v", s.ID, s.Title, s.Content, s.Created)
}

func (app *application) adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files := []string{
			"../ui/html/show.page.tmpl",
			"../ui/html/base.layout.tmpl",
			"../ui/html/footer.partial.tmpl",
			"../ui/html/login.tmpl",
		}
		// Парсинг файлов шаблонов...
		ts, err := template.ParseFiles(files...)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Error!")
		}

	}
	if r.Method == "POST" {
		login := r.FormValue("login")
		password := r.FormValue("password")
		_, err := app.article.Login(login, password)
		if err != nil {
			http.Redirect(w, r, "/admin", 404)
			return
		} else {
			http.Redirect(w, r, "/admin/menu", 301)
			return
		}
	}

}

func (app *application) adminMenu(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"../ui/html/show.page.tmpl",
		"../ui/html/base.layout.tmpl",
		"../ui/html/footer.partial.tmpl",
		"../ui/html/admin.panel.tmpl",
	}
	// Парсинг файлов шаблонов...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Fprintf(w, "Error!")
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		id := r.FormValue("id")

		if title != "" && content != "" {
			_, err := app.article.Insert(title, content)
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, "Empty title or content")
		}
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Fprintln(w, "Not a number!")
				return
			}
			err = app.article.Remove(idInt)
			if err != nil {
				fmt.Fprintln(w, "Not found Atricle with this id!")
				return
			}
		}
	}
}
