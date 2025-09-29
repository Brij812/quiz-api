package api

import (
	"encoding/json"
	"net/http"

	"github.com/brij812/quiz-api/internal/service"
)

type createQuizRequest struct {
	Title string `json:"title"`
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var req createQuizRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	quiz := service.CreateQuiz(req.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quiz)
}

func ListQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes := service.ListQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
