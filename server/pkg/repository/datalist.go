package repository

import (
	"database/sql"
	"log"
	"server/shemes"
)

type DataListSqlite struct {
	db *sql.DB
}

func NewDataListSqlite(db *sql.DB) *DataListSqlite {
	return &DataListSqlite{db: db}
}

func (dl *DataListSqlite) GetKey(num int) ([]shemes.QuizesAnswer, error) {
	var sliseQuiz []shemes.QuizesAnswer
	var quisTmp shemes.QuizesAnswer

	rows, err := dl.db.Query("SELECT id, text, value_answer, COALESCE(cause, ''), created, question_id FROM quizes_answer")
	if err != nil {
		log.Println(err)
		return sliseQuiz, err
	}
	for rows.Next() {
		if num == 0 {
			break
		}
		num--

		err = rows.Scan(&quisTmp.Id, &quisTmp.Text, &quisTmp.ValueAnswer, &quisTmp.Cause, &quisTmp.Created, &quisTmp.QuestionID)
		if err != nil {
			log.Println(err)
			return sliseQuiz, err
		}
		sliseQuiz = append(sliseQuiz, quisTmp)

	}

	return sliseQuiz, nil
}
