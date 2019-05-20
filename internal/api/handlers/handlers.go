package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/sterlingdeng/test_scores/internal/collections"
)

func GetList(col c.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := col.GetList()

		j, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		w.Write(j)

	}
}

func GetByID(col c.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var res *c.AggregatedExams

		switch v := col.(type) {

		case *c.StudentCollection:
			examNumber := mux.Vars(r)["id"]
			res = v.GetByID(examNumber)

		case *c.ExamCollection:
			studentID := mux.Vars(r)["number"]
			res = v.GetByID(studentID)
		}

		if res == nil {
			w.Write(nil)
		}

		j, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(j)
	}
}
