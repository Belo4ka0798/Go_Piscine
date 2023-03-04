package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type User struct {
	ID       int
	Login    string
	Password string
	Role     string
}
