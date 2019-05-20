package collections

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/r3labs/sse"
)

// ==================================
// ******** Initial Setup ***********
// ==================================
var mockExamsSlice = []*ExamData{
	&ExamData{
		Exam:      1,
		StudentID: "s1",
		Score:     50.00,
	},
	&ExamData{
		Exam:      2,
		StudentID: "s2",
		Score:     100.00,
	},
	&ExamData{
		Exam:      3,
		StudentID: "s3",
		Score:     75.00,
	},
	&ExamData{
		Exam:      3,
		StudentID: "s2",
		Score:     33.33,
	},
}

var studentCollection = &StudentCollection{
	Data: make(map[string][]*ExamData),
}

var examCollection = &ExamCollection{
	Data: make(map[int32][]*ExamData),
}

func init() {
	for _, exam := range mockExamsSlice {
		studentCollection.addData(exam)
		examCollection.addData(exam)
	}
}

// ==================================
// ******** Begin Testing ***********
// ==================================

func Test_HandleStream(t *testing.T) {
	ch := make(chan *sse.Event)
	s := &StudentCollection{Data: make(map[string][]*ExamData)}
	e := &ExamCollection{Data: make(map[int32][]*ExamData)}

	go HandleStream(ch, s, e)

	for _, examData := range mockExamsSlice {
		j, err := json.Marshal(examData)
		if err != nil {
			panic(err)
		}
		sseEvent := &sse.Event{
			ID:    nil,
			Data:  j,
			Event: nil,
			Retry: nil,
		}
		ch <- sseEvent
		ms := time.Millisecond
		time.Sleep(25 * ms)
	}

	if length := len(s.Data); length != 3 {
		t.Errorf("Could not add data to StudentCollection via streams. Got length %v, want length %v", length, 3)
	}

	if length := len(e.Data); length != 3 {
		t.Errorf("Could not add data to ExamCollection via streams. Got length %v, want length %v", length, 3)
	}
}

func Test_calculateAverageExamScore(t *testing.T) {
	type args struct {
		exams []*ExamData
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "test 1",
			args: args{
				exams: mockExamsSlice,
			},
			want: 64.582504,
		},
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
	tests := []struct {
		name      string
		studentID string
		want      bool
	}{
		{
			name:      "test1",
			studentID: "s1",
			want:      true,
		},
		{
			name:      "test2",
			studentID: "s2",
			want:      true,
		}, {
			name:      "test3",
			studentID: "s0",
			want:      false,
		}, {
			name:      "test4",
			studentID: "s3",
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, ok := studentCollection.Data[tt.studentID]; ok != tt.want {
				t.Errorf("Looking up %v in StudentCollection.Data,, want %v, got %v", tt.studentID, tt.want, ok)
			}
		})
	}
}

func TestStudentCollection_GetList(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test1",
			want: []string{"s1", "s2", "s3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := studentCollection.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentCollection.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExamCollection_GetByID(t *testing.T) {
	tests := []struct {
		name      string
		studentID string
		want      *AggregatedExams
	}{
		{
			name:      "test1",
			studentID: "s10",
			want:      nil,
		},
		{
			name:      "test2",
			studentID: "s1",
			want: &AggregatedExams{
				Exams:   []*ExamData{mockExamsSlice[0]},
				Average: 50.00,
			},
		},
		{
			name:      "test3",
			studentID: "s2",
			want: &AggregatedExams{
				Exams:   []*ExamData{mockExamsSlice[1], mockExamsSlice[3]},
				Average: (100.00 + 33.33) / 2.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := studentCollection.GetByID(tt.studentID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExamCollection.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
