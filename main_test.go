package main

import (
	_ "net/http"
	_ "net/http/httptest"
	"testing"

	_ "github.com/sterlingdeng/test_scores/internal/api/handlers"
	c "github.com/sterlingdeng/test_scores/internal/collections"
)

var mockExamsSlice = []*c.ExamData{
	&c.ExamData{
		Exam:      1,
		StudentID: "s1",
		Score:     50.00,
	},
	&c.ExamData{
		Exam:      2,
		StudentID: "s2",
		Score:     100.00,
	},
	&c.ExamData{
		Exam:      3,
		StudentID: "s3",
		Score:     75.00,
	},
	&c.ExamData{
		Exam:      3,
		StudentID: "s2",
		Score:     33.33,
	},
}

var studentCollection = c.StudentCollection{
	Data: make(map[string][]*c.ExamData),
}

var examCollection = c.ExamCollection{
	Data: make(map[int32][]*c.ExamData),
}

func init() {
	for _, exam := range mockExamsSlice {
		studentCollection.addData(exam)
		examCollection.addData(exam)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

// func TestStudentGetListHandler(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/students", nil)
// 	if err != nil {
// 		t.Errorf("error at creating NewRequest. Got %v, want nil", err)
// 	}

// 	rr := httptest.NewRecorder()
// }
