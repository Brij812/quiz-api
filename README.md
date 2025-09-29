Quiz API 🎯

A simple backend API for creating quizzes, adding questions, and submitting answers, built with Go + Chi.

🚀 Features

Create quizzes with titles

Add questions with multiple options and a correct answer

Fetch quiz questions (without revealing correct answers)

Submit answers and get a score

JSON error responses

Health check endpoint

In-memory store (easy to run, no DB setup)

Unit tests for scoring logic

🛠️ Tech Stack

Language: Go 1.25+

Router: Chi

Persistence: In-memory maps (future-ready for DB)

Testing: go test

⚡ Getting Started
1. Clone & Setup
git clone https://github.com/brij812/quiz-api.git
cd quiz-api
go mod tidy

2. Run the Server
go run ./cmd/server


Server starts on:

http://localhost:8080

3. Run Tests
go test ./...

📖 API Reference
Health Check
GET /health


Response

{ "status": "ok" }

Create Quiz
POST /quizzes/


Body

{ "title": "Go Basics" }


Response

{ "id": 1, "title": "Go Basics" }

List Quizzes
GET /quizzes/


Response

[
  { "id": 1, "title": "Go Basics" }
]

Add Question
POST /quizzes/{quizID}/questions


Body

{
  "text": "What is Go?",
  "options": ["Language", "Framework"],
  "correct_option_id": 1
}


Response

{
  "id": 1,
  "quiz_id": 1,
  "text": "What is Go?",
  "options": [
    { "id": 1, "question_id": 1, "text": "Language" },
    { "id": 2, "question_id": 1, "text": "Framework" }
  ]
}

Get Quiz Questions
GET /quizzes/{quizID}/questions


Response

[
  {
    "id": 1,
    "quiz_id": 1,
    "text": "What is Go?",
    "options": [
      { "id": 1, "question_id": 1, "text": "Language" },
      { "id": 2, "question_id": 1, "text": "Framework" }
    ]
  }
]


✅ Correct answer is hidden from response.

Submit Answers
POST /quizzes/{quizID}/submit


Body

{
  "answers": [
    { "question_id": 1, "option_id": 1 }
  ]
}


Response

{ "score": 1, "total": 1 }

🔥 Error Responses

All errors return JSON format with status code:

{ "error": "quiz not found" }

🧩 Future Improvements

Swap in-memory store with SQLite/Postgres

Support multiple question types (single choice, multiple choice, text with word limits)

Add authentication for quiz creators

Swagger/OpenAPI documentation