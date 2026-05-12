package task

import "time"


var nextID int64 = 1


type Task struct {
	ID          int64
	Description string
	Completed   bool
	CreatedBy   string
	Deadline    time.Time
}


func New(description string, createdBy string) Task {
	id := nextID
	nextID++

	return Task{
		ID:          id,
		Description: description,
		Completed:   false,
		CreatedBy:   createdBy,
		Deadline:    time.Now(),
	}
}

func (t *Task) Complete() {
	t.Completed = true
}


func (t *Task) SetDeadline(deadline time.Time) {
	t.Deadline = deadline
}


func (t *Task) SetDescription(description string) {
	t.Description = description
}
