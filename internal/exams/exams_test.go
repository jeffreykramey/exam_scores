package exams

import (
	"sync"
	"testing"
)

func TestCreateExam(t *testing.T) {
	examId := 1
	exam := CreateExam(examId)
	if exam == nil {
		t.Errorf("Expected non-nil ExamData pointer, but got nil")
	}
	if len(ExamsMap) != 1 {
		t.Errorf("Expected ExamsMap to have length 1, but got %d", len(ExamsMap))
	}

	// Set the original exam's AverageScore to test against after creating a dupeExam
	exam.AverageScore = 77.7
	// Attempt to create an exam with an existing ID
	dupeExam := CreateExam(examId)
	if dupeExam == nil {
		t.Errorf("Expected non-nil ExamData pointer, but got nil")
	}
	if dupeExam != exam {
		t.Errorf("Expected pointers to be the same")
	}
	// ExamsMap length shouldn't change
	if len(ExamsMap) != 1 {
		t.Errorf("Expected ExamsMap to have length 1, but got %d", len(ExamsMap))
	}

	// Create a new exam with a different ID
	exam = CreateExam(2)
	if exam == nil {
		t.Errorf("Expected non-nil ExamData pointer, but got nil")
	}
	if len(ExamsMap) != 2 {
		t.Errorf("Expected ExamsMap to have length 2, but got %d", len(ExamsMap))
	}
}

func TestAddScore(t *testing.T) {
	exam := CreateExam(123)
	exam.AddScore("student1", 80.0)
	if len(exam.Scores) != 1 {
		t.Errorf("Expected Scores length to be 1, but got %d", len(exam.Scores))
	}
	if exam.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 80.0, but got %f", exam.AverageScore)
	}

	// Add multiple scores to an exam
	exam.AddScore("student2", 89.5)
	exam.AddScore("student3", 70.5)
	if len(exam.Scores) != 3 {
		t.Errorf("Expected Scores length to be 3, but got %d", len(exam.Scores))
	}
	if exam.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 80.0, but got %f", exam.AverageScore)
	}

	// Add scores to different exams
	exam2 := CreateExam(2)
	exam2.AddScore("student1", 85.0)
	if len(exam2.Scores) != 1 {
		t.Errorf("Expected Scores length to be 1, but got %d", len(exam2.Scores))
	}
	if exam2.AverageScore != 85.0 {
		t.Errorf("Expected AverageScore to be 85.0, but got %f", exam2.AverageScore)
	}
}

func TestGetExam(t *testing.T) {
	exam := CreateExam(1)

	retrievedExam, err := GetExam(1)

	if retrievedExam == nil {
		t.Errorf("Expected retrieved exam to not be nil")
	}

	// Ensure that the error returned is nil
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	// Ensure that the retrieved exam is the same as the one initially created
	if retrievedExam != exam {
		t.Errorf("Expected pointers to be the same")
	}

	// Ensure that attempting to retrieve a non-existent exam returns an error
	_, err = GetExam(2)

	if err != ExamNotFoundError {
		t.Errorf("Expected error to be ExamNotFoundError, but got: %v", err)
	}
}

func TestGetExams(t *testing.T) {
	// Populate ExamsMap
	CreateExam(1)
	CreateExam(2)
	CreateExam(3)

	exams := GetExams()

	// Ensure that three exams are present in the list
	if len(exams) != 3 {
		t.Errorf("Expected 3 exams, got %d", len(exams))
	}
}

func TestUpdateAverageScore(t *testing.T) {
	exam := ExamData{Id: 1, Scores: []examScore{{"Alice", 80}, {"Bob", 90}},
		scoreLock: &sync.RWMutex{}}
	exam.updateAverageScore()

	// Ensure average is set correctly
	if exam.AverageScore != 85 {
		t.Errorf("Expected AverageScore to be 85, got %f", exam.AverageScore)
	}

	// Ensure average is set correctly after adding another score
	exam.Scores = append(exam.Scores, examScore{"Charlie", 70})
	exam.updateAverageScore()

	if exam.AverageScore != 80 {
		t.Errorf("Expected AverageScore to be 80, got %f", exam.AverageScore)
	}
}
