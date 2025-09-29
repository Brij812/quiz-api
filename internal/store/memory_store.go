package store

import "github.com/brij812/quiz-api/internal/models"

var (
	Quizzes         = make(map[int]models.Quiz)
	Questions       = make(map[int]models.Question)
	Options         = make(map[int]models.Option)
	QuizCounter     = 1
	QuestionCounter = 1
	OptionCounter   = 1
)
