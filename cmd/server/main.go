package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/brij812/quiz-api/internal/api"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	r.Route("/quizzes", func(r chi.Router) {
		r.Post("/", api.CreateQuiz)
		r.Get("/", api.ListQuizzes)
		r.Route("/{quizID}", func(r chi.Router) {
			r.Post("/questions", api.AddQuestion)
			r.Get("/questions", api.GetQuestions)
			r.Post("/submit", api.SubmitAnswers)
		})
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
