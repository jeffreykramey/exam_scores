package internal

import (
	"encoding/json"
	"exam_scores/internal/exams"
	"exam_scores/internal/students"
	"github.com/r3labs/sse/v2"
	"log"
)

type ScoreResponse struct {
	ExamID    int     `json:"exam"`
	StudentID string  `json:"studentId"`
	Score     float64 `json:"score"`
}

func ListenForScores() {
	scoreChannel := make(chan *sse.Event)
	client := sse.NewClient("http://live-test-scores.herokuapp.com/scores")
	err := client.SubscribeChan("messages", scoreChannel)
	if err != nil {
		return
	}

	for score := range scoreChannel {
		var response ScoreResponse
		err := json.Unmarshal(score.Data, &response)
		if err != nil {
			log.Println(err)
			continue
		}

		student, err := students.GetStudent(response.StudentID)
		if err == students.StudentNotFoundError {
			student = students.CreateStudent(response.StudentID)
		}
		student.AddExamScore(response.ExamID, response.Score)

		exam, err := exams.GetExam(response.ExamID)
		if err == exams.ExamNotFoundError {
			exam = exams.CreateExam(response.ExamID)
		}
		exam.AddScore(student.Id, response.Score)
	}
}
