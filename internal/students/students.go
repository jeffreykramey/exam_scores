package students

import (
	"errors"
	"sync"
)

var (
	StudentsMap  = make(map[string]*StudentData) // studentData indexed by student ID
	studentsLock sync.RWMutex
)

type StudentData struct {
	Id           string          `json:"id"`
	ExamScores   map[int]float64 `json:"exam_scores"`
	AverageScore float64         `json:"average_score"`
	scoreLock    *sync.RWMutex
}

var StudentNotFoundError = errors.New("student not found")

func CreateStudent(studentId string) *StudentData {
	studentsWriteLock()

	studentsLock.RLock()
	student, ok := StudentsMap[studentId]
	studentsLock.RUnlock()

	// Only create if studentId doesn't exist in map to avoid accidental overwrit
	if !ok {
		student = &StudentData{studentId, make(map[int]float64), 0.0,
			&sync.RWMutex{}}
		StudentsMap[studentId] = student
	}

	return student
}

func GetStudent(studentId string) (*StudentData, error) {
	studentsReadLock()

	student, ok := StudentsMap[studentId]

	if !ok {
		return nil, StudentNotFoundError
	}
	return student, nil
}

func GetStudents() []string {
	studentsReadLock()
	var students []string
	for _, student := range StudentsMap {
		students = append(students, student.Id)
	}

	return students
}

func (student *StudentData) AddExamScore(examId int, score float64) {
	studentsWriteLock()
	student.writeLock()
	student.ExamScores[examId] = score
	student.updateAverageScore()
}

func (student *StudentData) updateAverageScore() {
	sum := 0.0
	studentsLock.RLock()
	student.scoreLock.RLock()
	for _, score := range student.ExamScores {
		sum += score
	}
	student.scoreLock.RUnlock()
	studentsLock.RUnlock()

	studentsWriteLock()
	student.AverageScore = sum / float64(len(student.ExamScores))
}

func (student *StudentData) writeLock() {
	student.scoreLock.Lock()
	student.scoreLock.Unlock()
}

func (student *StudentData) readLock() {
	student.scoreLock.RLock()
	student.scoreLock.RUnlock()
}

func studentsWriteLock() {
	studentsLock.Lock()
	defer studentsLock.Unlock()
}

func studentsReadLock() {
	studentsLock.RLock()
	defer studentsLock.RUnlock()
}
