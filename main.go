package main

import (
	"exam_scores/internal"
	"exam_scores/internal/exams"
	"exam_scores/internal/students"
	"fmt"
	"net/http"
)

func main() {
	go internal.ListenForScores()
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/students", students.StudentsHandler)
	http.HandleFunc("/students/", students.StudentHandler)
	http.HandleFunc("/exams", exams.ExamsHandler)
	http.HandleFunc("/exams/", exams.ExamHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
