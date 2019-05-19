package student_handlers 

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/sterlingdeng/test_scores/internal/collections"
)

// GetAllStudents returns handler to retrieve all students that has taken an exam
func GetAllStudents(s *c.StudentCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students := s.GetAllStudents()

		j, err := json.Marshal(students)
		if err != nil {
			panic(err)	
		}

		w.Write(j)
	}
}

// GetStudentByID returns handler to retrieve student exam data by student ID
func GetStudentByID(s *c.StudentCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentID := mux.Vars(r)["id"]
		studentExamData := s.GetStudentById(studentID)

		if studentExamData == nil {
			w.Write(nil)
		}

		j, err := json.Marshal(studentExamData)
		if err != nil {
			panic(err)
		}

		w.Write(j)
	}
}