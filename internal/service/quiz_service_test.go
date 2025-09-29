package service

import (
	"testing"

	"github.com/brij812/quiz-api/internal/models"
)

func TestSubmitAnswers_Empty(t *testing.T) {
	CreateQuiz("Go Basics")
	AddQuestion(1, "What is Go?", []string{"Language", "Framework"}, 1)

	answers := []models.Answer{}
	score, total, _ := SubmitAnswers(1, answers)
	if score != 0 || total != 1 {
		t.Errorf("expected score=0, total=1, got score=%d, total=%d", score, total)
	}
}

func TestAddQuestion_InvalidText(t *testing.T) {
	_, err := AddQuestion(999, "", []string{"Yes", "No"}, 1)
	if err == nil {
		t.Errorf("expected error for empty question text, got nil")
	}
}

func TestAddQuestion_InvalidCorrectOption(t *testing.T) {
	CreateQuiz("Invalid Test")
	_, err := AddQuestion(2, "Test?", []string{"A", "B"}, 999)
	if err == nil {
		t.Errorf("expected error for invalid correct_option_id, got nil")
	}
}
