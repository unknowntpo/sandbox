package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type createTaskPayload struct {
	rID     string
	title   string
	content string
}

// POST /task
func createTaskHandler(c *gin.Context) {
	rID := "1234"
	redirURL := fmt.Sprintf("/queue/%s", rID)

	q.Enqueue(createTaskPayload{
		rID:     rID,
		title:   "hello",
		content: "first task",
	})

	c.JSON(200, gin.H{
		"req_id": rID,
		"url":    redirURL,
	})
}

const (
	StatusPending = iota
	StatusProcessing
	StatusDone
	StatusError
)

// continue consuming q, do create task.
func doCreateTask(q chan createTaskPayload) {
	for {
		payload <- q

	}
}

// /queue/:req_id
func requestStatusHandler(c *gin.Context) {

}

func main() {
	go startQueue()
	Start()
}

var q chan createTaskPayload

func startQueue(chan createTaskPayload) {
	go doCreateTask()

}

func Start() {
	r := gin.Default()
	r.POST("/task", createTaskHandler)
	r.GET("/queue/:req_id", requestStatusHandler)
	r.Run()
}
