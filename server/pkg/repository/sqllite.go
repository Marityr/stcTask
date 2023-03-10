package repository

import (
	"database/sql"
	"log"
	"server/shemes"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlliteDB() (*sql.DB, error) {
	var once sync.Once

	DB, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("cannot open an SQLite memory database: %v", err)
		return DB, err
	}

	onse := func() {
		createTableSQL := `CREATE TABLE "quizes_answer" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "text" varchar(200) NOT NULL, "value_answer" varchar(200) NOT NULL, "cause" varchar(200) NULL, "created" datetime NOT NULL, "question_id" integer NOT NULL);`
		log.Println("Create table...")
		statement, err := DB.Prepare(createTableSQL)
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
		log.Println("table created")

		err = RewriteDB(DB)
		if err != nil {
			log.Fatalf("not old")
		}
	}

	go once.Do(onse)

	return DB, nil
}

func RewriteDB(db *sql.DB) error {
	var sliseQuiz []shemes.QuizesAnswer
	var quisTmp shemes.QuizesAnswer
	oldDB, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatalf("cannot open an SQLite memory database: %v", err)
		return err
	}
	defer oldDB.Close()

	rows, err := oldDB.Query("SELECT id, text, value_answer, COALESCE(cause, ''), created, question_id FROM quizes_answer")
	if err != nil {
		log.Fatal(err)
		return err
	}
	for rows.Next() {
		err = rows.Scan(&quisTmp.Id, &quisTmp.Text, &quisTmp.ValueAnswer, &quisTmp.Cause, &quisTmp.Created, &quisTmp.QuestionID)
		if err != nil {
			log.Fatal(err)
			return err
		}
		sliseQuiz = append(sliseQuiz, quisTmp)
	}

	for _, data := range sliseQuiz {
		insertSQL := `INSERT INTO quizes_answer(id, text, value_answer, cause, created, question_id) VALUES (?, ?, ?, ?, ?, ?)`
		statement, err := db.Prepare(insertSQL)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec(data.Id, data.Text, data.ValueAnswer, data.Cause, data.Created, data.QuestionID)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	log.Println("rewrite end")
	return nil
}
