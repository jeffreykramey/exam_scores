package exams

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExamHandler(t *testing.T) {
	var tests = []struct {
		param        string
		expectedCode int
	}{
		{"test_one", http.StatusBadRequest},
		{"123", http.StatusNotFound},
		{"1", http.StatusOK},
	}

	CreateExam(1)

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.param)
		t.Run(testName, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			request, err := http.NewRequest("GET", fmt.Sprintf("/exams/%s", tt.param), nil)
			if err != nil {
				t.Fatal(err)
			}
			ExamHandler(responseRecorder, request)
			if status := responseRecorder.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v wanted %v", status, tt.expectedCode)
			}
		})
	}
}

func TestExamsHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/exams", nil)
	if err != nil {
		t.Fatal(err)
	}

	CreateExam(1)
	ExamsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `[1]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", rr.Body.String(), expected)
	}
}
