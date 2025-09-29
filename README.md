# Quiz API 🎯

A simple backend API for creating quizzes, adding questions, and submitting answers — built with **Go + Chi**.

---

## 🚀 Features

- Create quizzes with titles  
- Add questions with multiple options and a correct answer  
- Fetch quiz questions (**without revealing correct answers**)  
- Submit answers and get a score  
- JSON error responses for consistency  
- Health check endpoint (`/health`)  
- In-memory store (no database setup needed)  
- Unit tests for core scoring logic  

---

## 🛠️ Tech Stack

| Component   | Technology      |
|-------------|-----------------|
| Language    | Go 1.25+        |
| Router      | [Chi](https://github.com/go-chi/chi) |
| Persistence | In-memory maps  |
| Testing     | `go test`       |

---

## ⚡ Getting Started

### 1. Clone & Setup
```bash
git clone https://github.com/brij812/quiz-api.git
cd quiz-api
go mod tidy
```

### 2. Run the Server
```bash
go run ./cmd/server
```

Server starts at:
```
http://localhost:8080
```

### 3. Run Tests
```bash
go test ./...
```

---

## 📖 API Reference

### 🔹 Health Check
**Request**
```http
GET /health
```
**Response**
```json
{ "status": "ok" }
```

---

### 🔹 Create Quiz
**Request**
```http
POST /quizzes/
Content-Type: application/json

{
  "title": "Go Basics"
}
```

**Response**
```json
{ "id": 1, "title": "Go Basics" }
```

---

### 🔹 List Quizzes
**Request**
```http
GET /quizzes/
```

**Response**
```json
[
  { "id": 1, "title": "Go Basics" }
]
```

---

### 🔹 Add Question
**Request**
```http
POST /quizzes/{quizID}/questions
Content-Type: application/json

{
  "text": "What is Go?",
  "options": ["Language", "Framework"],
  "correct_option_id": 1
}
```

**Response**
```json
{
  "id": 1,
  "quiz_id": 1,
  "text": "What is Go?",
  "options": [
    { "id": 1, "question_id": 1, "text": "Language" },
    { "id": 2, "question_id": 1, "text": "Framework" }
  ]
}
```

---

### 🔹 Get Quiz Questions
**Request**
```http
GET /quizzes/{quizID}/questions
```

**Response**
```json
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
```

✅ The correct answer is **hidden**.

---

### 🔹 Submit Answers
**Request**
```http
POST /quizzes/{quizID}/submit
Content-Type: application/json

{
  "answers": [
    { "question_id": 1, "option_id": 1 }
  ]
}
```

**Response**
```json
{ "score": 1, "total": 1 }
```

---

## 🔥 Error Handling

All errors return JSON format with proper status codes:
```json
{ "error": "quiz not found" }
```

---

## 🧩 Future Improvements

- Replace in-memory store with SQLite/Postgres  
- Support multiple question types (single-choice, multiple-choice, text with word limits)  
- Add authentication for quiz creators  
- Swagger/OpenAPI documentation  

---

## 📜 License

This project is released under the [MIT License](LICENSE).
