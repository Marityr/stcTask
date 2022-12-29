package repository

import (
	"database/sql"
	"stcTask/server/shemes"
)

type DataList interface {
	GetKey(num int) ([]shemes.QuizesAnswer, error)
}

type Repository struct {
	DataList
}

func NewReposiroty(db *sql.DB) *Repository {
	return &Repository{
		DataList: NewDataListSqlite(db),
	}
}
