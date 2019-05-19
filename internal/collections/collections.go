package collections 

import (
	"encoding/json"
	"sync"
	// "fmt"
	"github.com/r3labs/sse"
)

// ExamData provides struct for msg.Data to unmarshall into
type ExamData struct {
	Exam      int32   `json:"exam"`
	StudentID string  `json:"studentId"`
	Score     float32 `json:"score"`
}

// HandleStream function is used to take in take in stream data and parse the information
func HandleStream(ch chan *sse.Event, s *StudentCollection, e *ExamCollection) {
	for {
		msg := <-ch // channels are blocking.. so even though its in a forever for loop, it will block until there is something in the channel

		var examData ExamData

		err := json.Unmarshal(msg.Data, &examData)
		if err != nil {
			panic(err)
		}

		s.addData(&examData)
		e.addData(&examData)
	}
}

// StudentCollection provides structure to hold student exam data
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

// GetAllStudents returns all the students, in the form of an 
// slice of strings, that have taken an exam
// Used for the /students api endpoint
func (s *StudentCollection) GetAllStudents() []string {
	studentList := []string{}
	for studentID := range s.Data {
		studentList = append(studentList, studentID)
	}
	return studentList
}

type StudentExamData struct {
	Exams []*ExamData `json:"exams"` 
	Average float32 `json:"average"`
}

func (s *StudentCollection) GetStudentById(id string) *StudentExamData {
	studentExamData := &StudentExamData{
		Exams: []*ExamData{},
	}

	exams, ok := s.Data[id]

	if !ok {
		return nil
	}

	studentExamData.Exams = exams 
	studentExamData.Average = calculateAverageExam(exams)	

	return studentExamData
}

func calculateAverageExam(exams []*ExamData) float32 {
	count := float32(len(exams))
	var sum float32
	for _, exam := range exams {
		sum += exam.Score
	}
	return sum/count
}

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

// GetAllExams gets the list of all exams, as a slice of int32
func (e *ExamCollection) GetAllExams() []int32 {
	examList := []int32{}
	for examID := range e.Data {
		examList = append(examList, examID)
	}
	return examList
}

type ExamStatsData struct {
	Exams []*ExamData `json:"exams"`
	Average float32 `json:"average"`
}

func (e *ExamCollection) GetExamByID(id int32) *ExamStatsData {
	examStatsData := &ExamStatsData{
		Exams: []*ExamData{},
	}

	exams, ok := e.Data[id]
	if !ok {
		return nil
	}

	examStatsData.Exams = exams
	examStatsData.Average = calculateAverageExam(exams)

	return examStatsData
}

