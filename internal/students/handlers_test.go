package students

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStudentHandler(t *testing.T) {
	var tests = []struct {
		param        string
		expectedCode int
	}{
		{"john_doe", http.StatusNotFound},
		{"jane_foster", http.StatusOK},
	}

	CreateStudent("jane_foster")

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.param)
		t.Run(testName, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			request, err := http.NewRequest("GET", fmt.Sprintf("/students/%s", tt.param), nil)
			if err != nil {
				t.Fatal(err)
			}
			StudentHandler(responseRecorder, request)
			if status := responseRecorder.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v wanted %v", status, tt.expectedCode)
			}
		})
	}
}

func TestStudentsHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/students", nil)
	if err != nil {
		t.Fatal(err)
	}

	CreateStudent("jane_foster")
	StudentsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `["jane_foster"]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", rr.Body.String(), expected)
	}
}
