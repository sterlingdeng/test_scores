package collections

import (
	"sync"
	"testing"
)

func Test_calculateAverageExamScore(t *testing.T) {
	type args struct {
		exams []*ExamData
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAverageExamScore(tt.args.exams); got != tt.want {
				t.Errorf("calculateAverageExamScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentCollection_addData(t *testing.T) {
	type fields struct {
		Data map[string][]*ExamData
		mu   sync.Mutex
	}
	type args struct {
		examData *ExamData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentCollection{
				Data: tt.fields.Data,
				mu:   tt.fields.mu,
			}
			s.addData(tt.args.examData)
		})
	}
}

func TestExamCollection_addData(t *testing.T) {
	type fields struct {
		Data map[int32][]*ExamData
		mu   sync.Mutex
	}
	type args struct {
		examData *ExamData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExamCollection{
				Data: tt.fields.Data,
				mu:   tt.fields.mu,
			}
			e.addData(tt.args.examData)
		})
	}
}
