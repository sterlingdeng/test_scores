package collections

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/r3labs/sse"
)

// ExamData provides a struct for msg.Data to unmarshal into
type ExamData struct {
	Exam      int32   `json:"exam"`
	StudentID string  `json:"studentId"`
	Score     float32 `json:"score"`
}

// HandleStream function is used to take in take in stream data and parse the information
func HandleStream(ch chan *sse.Event, s *StudentCollection, e *ExamCollection) {
	for {
		var examData ExamData
		msg := <-ch

		err := json.Unmarshal(msg.Data, &examData)
		if err != nil {
			panic(err)
		}
		s.addData(&examData)
		e.addData(&examData)
	}
}

// AggregatedExams is the response for requests that hit the GET endpoints
type AggregatedExams struct {
	Exams   []*ExamData `json:"exams"`
	Average float32     `json:"average"`
}

func calculateAverageExamScore(exams []*ExamData) float32 {
	count := float32(len(exams))
	var sum float32
	for _, exam := range exams {
		sum += exam.Score
	}
	return sum / count
}

// Collection defines the interface which is used by
// the REST API handler
type Collection interface {
	GetList() []string
	GetByID(string) *AggregatedExams
}

// StudentCollection provides quick look up of Exam Data by
// student ID
type StudentCollection struct {
	Data map[string][]*ExamData
	mu   sync.Mutex
}

func (s *StudentCollection) addData(examData *ExamData) {
	s.mu.Lock()
	defer s.mu.Unlock()

	studentID := examData.StudentID
	s.Data[studentID] = append(s.Data[studentID], examData)
}

// GetList returns all the students, in the form of an
// slice of strings, that have taken an exam
// Used for the /students api endpoint
func (s *StudentCollection) GetList() []string {
	studentList := []string{}
	for studentID := range s.Data {
		studentList = append(studentList, studentID)
	}
	return studentList
}

// GetByID returns AggregatedExams, selected by student ID
func (s *StudentCollection) GetByID(id string) *AggregatedExams {
	studentExamData := &AggregatedExams{
		Exams: []*ExamData{},
	}

	exams, ok := s.Data[id]
	if !ok {
		return nil
	}

	studentExamData.Exams = exams
	studentExamData.Average = calculateAverageExamScore(exams)

	return studentExamData
}

// ExamCollection provides quick look up of Exam Data by
// exam ID
type ExamCollection struct {
	Data map[int32][]*ExamData
	mu   sync.Mutex
}

func (e *ExamCollection) addData(examData *ExamData) {
	e.mu.Lock()
	defer e.mu.Unlock()

	examID := examData.Exam

	e.Data[examID] = append(e.Data[examID], examData)
}

// GetList gets the list of all exams, as a slice of int32
func (e *ExamCollection) GetList() []string {
	examList := []string{}
	for examID := range e.Data {
		examList = append(examList, strconv.Itoa(int(examID)))
	}
	return examList
}

// GetByID returns AggregatedExams, selected by exam ID
func (e *ExamCollection) GetByID(id string) *AggregatedExams {
	examStatsData := &AggregatedExams{
		Exams: []*ExamData{},
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	exams, ok := e.Data[int32(intID)]
	if !ok {
		return nil
	}

	examStatsData.Exams = exams
	examStatsData.Average = calculateAverageExamScore(exams)

	return examStatsData
}
