package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID    string `json:"id"`
	Desc  string `json:"desc"`
	Topic string `json:"topic"`
	//Options []option
	// CorrectAnswer string
}

var questions = []question{
	{ID: "1", Desc: "First Question Description", Topic: "History"},
	{ID: "2", Desc: "Second Question Description", Topic: "Math"},
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func createQuestion(c *gin.Context) {
	var newQuestion question
	if err := c.BindJSON(&newQuestion); err != nil {
		return
	}
	questions = append(questions, newQuestion)
	c.IndentedJSON(http.StatusCreated, newQuestion)
}

func getQuestionById(id string) (*question, error) {
	for i, q := range questions {
		if q.ID == id {
			return &questions[i], nil
		}
	}
	return nil, errors.New("question not found")
}

func questionById(c *gin.Context) {
	id := c.Param("id")
	question, err := getQuestionById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, question)
}

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.GET("/question/:id", questionById)
	router.POST("/question", createQuestion)
	router.Run("localhost:8080")
}
