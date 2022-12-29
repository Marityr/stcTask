package shemes

import "time"

type QuizesAnswer struct {
	Id          int
	Text        string
	ValueAnswer string
	Cause       string
	Created     time.Time
	QuestionID  int
}
