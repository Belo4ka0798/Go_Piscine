package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"go_day06/ex01/pkg"
	"log"
	"strconv"
)

// ArticleModel - Определяем тип который обертывает пул подключения sql.DB
type ArticleModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *ArticleModel) Insert(title, content string) (int, error) {
	stmt := `INSERT INTO Articles (Title, Content, Created)
    VALUES($1, $2, current_date);`
	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}
	queryID := `select ID from Articles as a Where a.Title = $1 and a.Content = $2 `
	var id string
	err = m.DB.QueryRow(queryID, title, content).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	i, err := strconv.Atoi(id)
	fmt.Printf("Sucsess insert %v", i)
	return i, nil
}

func (m *ArticleModel) Remove(id int) error {
	stmt := `-- DELETE FROM public.atricles WHERE ID = $1;`
	_, err := m.DB.Exec(stmt, id)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *ArticleModel) Get(id int) (*models.Article, error) {
	stmt := `SELECT ID, Title, Content, Created FROM Articles
    WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Article{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// AllArticles - Метод возвращает Все статьи
func (m *ArticleModel) AllArticles() ([]*models.Article, error) {
	stmt := `SELECT id, title, content, created FROM Articles ORDER BY created`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*models.Article
	for rows.Next() {
		s := &models.Article{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
		if err != nil {
			return nil, err
		}
		articles = append(articles, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *ArticleModel) Login(login string, pass string) (*models.User, error) {
	stmt := `SELECT ID, Login, Password, Role FROM Users
    WHERE Login = $1 AND Password = $2`
	row := m.DB.QueryRow(stmt, login, pass)
	s := &models.User{}
	err := row.Scan(&s.ID, &s.Login, &s.Password, &s.Role)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s %s", s.Login, s.Role)
	return s, nil
}
