package students

import (
	"sync"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	studentId := "jane_foster"
	student := CreateStudent(studentId)
	if student == nil {
		t.Errorf("Expected non-nil pointer, but got nil")
	}
	if len(StudentsMap) != 1 {
		t.Errorf("Expected StudentssMap to have length 1, but got %d", len(StudentsMap))
	}

	// Set the original student's AverageScore to test against after creating attempting to create a duplicate student
	student.AverageScore = 88.8
	// Attempt to create n student with an existing ID
	dupeStudent := CreateStudent(studentId)
	if dupeStudent == nil {
		t.Errorf("Expected non-nil pointer, but got nil")
	}
	if dupeStudent != student {
		t.Errorf("Expected pointers to be the same")
	}
	// StudentsMap length shouldn't change
	if len(StudentsMap) != 1 {
		t.Errorf("Expected StudentsMap to have length 1, but got %d", len(StudentsMap))
	}

	// Create a new student with a different ID
	student = CreateStudent("john_doe")
	if student == nil {
		t.Errorf("Expected non-nil pointer, but got nil")
	}
	if len(StudentsMap) != 2 {
		t.Errorf("Expected StudentsMap to have length 2, but got %d", len(StudentsMap))
	}
}

func TestAddExamScore(t *testing.T) {
	student := CreateStudent("jane_foster")
	student.AddExamScore(1, 80.0)
	if len(student.ExamScores) != 1 {
		t.Errorf("Expected Scores length to be 1, but got %d", len(student.ExamScores))
	}
	if student.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 80.0, but got %f", student.AverageScore)
	}

	// Add multiple test scores for a student
	student.AddExamScore(2, 89.5)
	student.AddExamScore(3, 70.5)
	if len(student.ExamScores) != 3 {
		t.Errorf("Expected ExamScores length to be 3, but got %d", len(student.ExamScores))
	}
	if student.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 80.0, but got %f", student.AverageScore)
	}

	// Add scores to different exams for a new student
	student2 := CreateStudent("john_doe")
	student2.AddExamScore(4, 85.0)
	if len(student2.ExamScores) != 1 {
		t.Errorf("Expected Scores length to be 1, but got %d", len(student2.ExamScores))
	}
	if student2.AverageScore != 85.0 {
		t.Errorf("Expected AverageScore to be 85.0, but got %f", student2.AverageScore)
	}
}

func TestGetStudent(t *testing.T) {
	studentId := "jane_foster"
	student := CreateStudent(studentId)

	retrievedStudent, err := GetStudent(studentId)

	if retrievedStudent == nil {
		t.Errorf("Expected retrieved student to not be nil")
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}

	// Ensure that the retrieved student is the same as the one initially created
	if retrievedStudent != student {
		t.Errorf("Expected pointers to be the same")
	}

	// Ensure that attempting to retrieve a non-existent student returns an error
	_, err = GetStudent("doesn't_exist")

	if err != StudentNotFoundError {
		t.Errorf("Expected error to be StudentNotFoundError, but got: %v", err)
	}
}

func TestGetStudents(t *testing.T) {
	// Populate ExamsMap
	CreateStudent("jane_foster")
	CreateStudent("john_doe")
	CreateStudent("billy_mays")

	students := GetStudents()

	if len(students) != 3 {
		t.Errorf("Expected 3 students, got %d", len(students))
	}
}

func TestUpdateAverageScore(t *testing.T) {
	student := StudentData{Id: "jane_foster", ExamScores: map[int]float64{1: 75.5, 2: 84.5},
		scoreLock: &sync.RWMutex{}}
	student.updateAverageScore()

	// Ensure average is set correctly
	if student.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 85, got %f", student.AverageScore)
	}

	// Ensure average is set correctly after adding another score
	student.ExamScores[3] = 80.0
	student.updateAverageScore()

	if student.AverageScore != 80.0 {
		t.Errorf("Expected AverageScore to be 80, got %f", student.AverageScore)
	}
}
