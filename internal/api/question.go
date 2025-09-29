package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/brij812/quiz-api/internal/models"
	"github.com/brij812/quiz-api/internal/service"
)

type addQuestionRequest struct {
	Text            string   `json:"text"`
	Options         []string `json:"options"`
	CorrectOptionID int      `json:"correct_option_id"`
}

func AddQuestion(w http.ResponseWriter, r *http.Request) {
	quizID, _ := strconv.Atoi(chi.URLParam(r, "quizID"))
	var req addQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	q, err := service.AddQuestion(quizID, req.Text, req.Options, req.CorrectOptionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	quizID, _ := strconv.Atoi(chi.URLParam(r, "quizID"))
	qs, err := service.GetQuestions(quizID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(qs)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	quizID, _ := strconv.Atoi(chi.URLParam(r, "quizID"))
	var req models.SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	score, total, err := service.SubmitAnswers(quizID, req.Answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{
		"score": score,
		"total": total,
	})
}
