package model

import "database/sql"

type Car struct {
	Id       int           `json:"id" db:"id"`
	RegNum   string        `json:"regNum" db:"regnum"`
	Mark     string        `json:"mark" db:"mark"`
	Model    string        `json:"model" db:"model"`
	Year     int           `json:"year,omitempty" db:"-"`
	NullYear sql.NullInt16 `json:"-" db:"year"`
	Owner    struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic,omitempty"`
	} `json:"owner"`
	Name           string         `json:"-" db:"name"`
	Surname        string         `json:"-" db:"surname"`
	NullPatronymic sql.NullString `json:"-" db:"patronymic"`
	Patronymic     string         `json:"-" db:"-"`
}

func (f Car) IsEmpty() bool {
	return false
}
