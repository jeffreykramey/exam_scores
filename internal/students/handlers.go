package students

import (
	"encoding/json"
	"net/http"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Path[len("/students/"):]

	student, err := GetStudent(studentID)

	if err == StudentNotFoundError {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	students := GetStudents()

	jsonBytes, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
