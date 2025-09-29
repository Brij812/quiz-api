package service

import (
	"errors"

	"github.com/brij812/quiz-api/internal/models"
	"github.com/brij812/quiz-api/internal/store"
)

func CreateQuiz(title string) (models.Quiz, error) {
	if len(title) == 0 {
		return models.Quiz{}, errors.New("quiz title cannot be empty")
	}
	quiz := models.Quiz{
		ID:    store.QuizCounter,
		Title: title,
	}
	store.Quizzes[store.QuizCounter] = quiz
	store.QuizCounter++
	return quiz, nil
}

func ListQuizzes() []models.Quiz {
	result := []models.Quiz{}
	for _, q := range store.Quizzes {
		result = append(result, q)
	}
	return result
}

func AddQuestion(quizID int, text string, options []string, correctID int) (models.Question, error) {
	if _, ok := store.Quizzes[quizID]; !ok {
		return models.Question{}, errors.New("quiz not found")
	}
	if len(text) == 0 {
		return models.Question{}, errors.New("question text cannot be empty")
	}
	if len(text) > 300 {
		return models.Question{}, errors.New("question text exceeds 300 characters")
	}
	if len(options) < 2 {
		return models.Question{}, errors.New("must have at least 2 options")
	}
	q := models.Question{
		ID:              store.QuestionCounter,
		QuizID:          quizID,
		Text:            text,
		Options:         []models.Option{},
		CorrectOptionID: correctID,
	}
	for _, opt := range options {
		o := models.Option{
			ID:         store.OptionCounter,
			QuestionID: q.ID,
			Text:       opt,
		}
		store.Options[store.OptionCounter] = o
		q.Options = append(q.Options, o)
		store.OptionCounter++
	}
	valid := false
	for _, o := range q.Options {
		if o.ID == correctID {
			valid = true
		}
	}
	if !valid {
		return models.Question{}, errors.New("invalid correct_option_id")
	}
	store.Questions[store.QuestionCounter] = q
	store.QuestionCounter++
	return q, nil
}

func GetQuestions(quizID int) ([]models.Question, error) {
	if _, ok := store.Quizzes[quizID]; !ok {
		return nil, errors.New("quiz not found")
	}
	result := []models.Question{}
	for _, q := range store.Questions {
		if q.QuizID == quizID {
			copyQ := q
			copyQ.CorrectOptionID = 0
			result = append(result, copyQ)
		}
	}
	return result, nil
}

func SubmitAnswers(quizID int, answers []models.Answer) (int, int, error) {
	if _, ok := store.Quizzes[quizID]; !ok {
		return 0, 0, errors.New("quiz not found")
	}
	questions := []models.Question{}
	for _, q := range store.Questions {
		if q.QuizID == quizID {
			questions = append(questions, q)
		}
	}
	total := len(questions)
	score := 0
	for _, ans := range answers {
		for _, q := range questions {
			if q.ID == ans.QuestionID && q.CorrectOptionID == ans.OptionID {
				score++
			}
		}
	}
	return score, total, nil
}
