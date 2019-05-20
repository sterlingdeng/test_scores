package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/sterlingdeng/test_scores/internal/collections"
)

// GetList gets a list of entry keys of a collection.
func GetList(col c.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := col.GetList()

		j, err := json.Marshal(list)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}

// GetByID returns the AggregatedData struct, selected by a particular ID
func GetByID(col c.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res *c.AggregatedExams

		switch v := col.(type) {
		case *c.StudentCollection:
			studentID := mux.Vars(r)["id"]
			res = v.GetByID(studentID)

		case *c.ExamCollection:
			examNumber := mux.Vars(r)["number"]
			res = v.GetByID(examNumber)
		}

		if res == nil {
			w.WriteHeader(400)
			w.Write(nil)
		}

		j, err := json.Marshal(res)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}
