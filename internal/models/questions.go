package models

type Question struct {
	ID              int      `json:"id"`
	QuizID          int      `json:"quiz_id"`
	Text            string   `json:"text"`
	Options         []Option `json:"options"`
	CorrectOptionID int      `json:"-"`
}

type Answer struct {
	QuestionID int `json:"question_id"`
	OptionID   int `json:"option_id"`
}

type SubmitRequest struct {
	Answers []Answer `json:"answers"`
}
