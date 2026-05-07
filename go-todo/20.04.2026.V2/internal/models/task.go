package models

type Task struct {
	Id          int
	Description string
	Completed   bool
}

var Tasks = make(map[int]*Task)
var NextID = 1
