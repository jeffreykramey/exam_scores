package exams

import (
	"errors"
	"sync"
)

var (
	ExamsMap  = make(map[int]*ExamData) // examData indexed by examId
	examsLock sync.RWMutex
)

type ExamData struct {
	Id           int         `json:"id"`
	Scores       []examScore `json:"scores"`
	AverageScore float64     `json:"average_score"`
	scoreLock    *sync.RWMutex
}

type examScore struct {
	StudentId string  `json:"student_id"`
	Score     float64 `json:"score"`
}

var ExamNotFoundError = errors.New("exam not found")

func CreateExam(examId int) *ExamData {
	examsWriteLock()

	examsLock.RLock()
	exam, ok := ExamsMap[examId]
	examsLock.RUnlock()

	// Only create if examId doesn't exist in map to avoid accidental overwrite
	if !ok {
		exam = &ExamData{examId, []examScore{}, 0.0, &sync.RWMutex{}}
		ExamsMap[examId] = exam
	}

	return exam
}

func (exam *ExamData) AddScore(studentId string, score float64) {
	examsWriteLock()
	exam.writeLock()
	exam.Scores = append(exam.Scores, examScore{studentId, score})
	exam.updateAverageScore()
}

func GetExam(examId int) (*ExamData, error) {
	examsReadLock()
	exam, ok := ExamsMap[examId]

	if !ok {
		return nil, ExamNotFoundError
	}
	return exam, nil
}

func GetExams() []int {
	examsReadLock()
	var exams []int
	for _, exam := range ExamsMap {
		exams = append(exams, exam.Id)
	}

	return exams
}

func (exam *ExamData) updateAverageScore() {
	sum := 0.0
	examsLock.RLock()
	exam.scoreLock.RLock()
	for _, entry := range exam.Scores {
		sum += entry.Score
	}
	exam.scoreLock.RUnlock()
	examsLock.RUnlock()

	examsWriteLock()
	exam.AverageScore = sum / float64(len(exam.Scores))
}

func (exam *ExamData) writeLock() {
	exam.scoreLock.Lock()
	exam.scoreLock.Unlock()
}

func (exam *ExamData) readLock() {
	exam.scoreLock.RLock()
	exam.scoreLock.RUnlock()
}

func examsWriteLock() {
	examsLock.Lock()
	defer examsLock.Unlock()
}

func examsReadLock() {
	examsLock.RLock()
	defer examsLock.RUnlock()
}
