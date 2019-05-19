package main

import (
	"net/http"

  c "github.com/sterlingdeng/test_scores/internal/collections"
  "github.com/sterlingdeng/test_scores/internal/api/student"
  "github.com/sterlingdeng/test_scores/internal/api/exam"

	"github.com/gorilla/mux"
	"github.com/r3labs/sse"
)

func main() {
	events := make(chan *sse.Event)
	s := &c.StudentCollection{Data: make(map[string][]*c.ExamData)}
	e := &c.ExamCollection{Data: make(map[int32][]*c.ExamData)}

	client := sse.NewClient("http://live-test-scores.herokuapp.com/scores")
	client.SubscribeChan("score", events)

	go c.HandleStream(events, s, e)

	r := mux.NewRouter()

	r.HandleFunc("/students", student_handlers.GetAllStudents(s)).Methods("GET")
	r.HandleFunc("/students/{id}", student_handlers.GetStudentByID(s)).Methods("GET")

	r.HandleFunc("/exams", exam_handlers.GetAllExams(e)).Methods("GET")
	r.HandleFunc("/exams/{number}", exam_handlers.GetExamByID(e)).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
