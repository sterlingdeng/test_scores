package exam_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	c "github.com/sterlingdeng/test_scores/internal/collections"
)

func GetAllExams(e *c.ExamCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exams := e.GetAllExams()

		j, err := json.Marshal(exams)
		if err != nil {
			panic(err)
		}
		
		w.Write(j)
	}
}

func GetExamByID(e *c.ExamCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		examID, err := strconv.Atoi(mux.Vars(r)["number"])
		if err != nil {
			panic(err)
		}

		examStatsData := e.GetExamByID(int32(examID))
		if examStatsData == nil {
			w.Write(nil)
		}

		j, err := json.Marshal(examStatsData) 
		if err != nil {
			panic(err)
		}

		w.Write(j)
	}
}