package exams

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ExamHandler(w http.ResponseWriter, r *http.Request) {
	examStr := r.URL.Path[len("/exams/"):]
	examId, err := strconv.Atoi(examStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exam, err := GetExam(examId)

	if err == ExamNotFoundError {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(exam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func ExamsHandler(w http.ResponseWriter, r *http.Request) {
	examsList := GetExams()

	jsonBytes, err := json.Marshal(examsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
