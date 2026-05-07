package repository

import "Todo/internal/models"

type TaskRepository struct {
	tasks  map[int]*models.Task
	nextID int
}


func NewTaskRepository() *TaskRepository {
	repo := &TaskRepository{
		tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
	

	repo.Add("Сделать домашнее задание")
	repo.Add("Пойти в магазин")
	repo.Add("Курс по Go")
	
	return repo
}

func (r *TaskRepository) GetAll() []*models.Task {
	var list []*models.Task
	for _, task := range r.tasks {
		list = append(list, task)
	}
	return list
}

func (r *TaskRepository) Add(description string) int {
	task := &models.Task{
		Id:          r.nextID,
		Description: description,
		Completed:   false,
	}
	r.tasks[r.nextID] = task
	r.nextID++
	return task.Id
}

func (r *TaskRepository) Delete(id int) bool {
	if _, exists := r.tasks[id]; exists {
		delete(r.tasks, id)
		return true
	}
	return false
}

func (r *TaskRepository) Complete(id int) bool {
	if task, exists := r.tasks[id]; exists {
		task.Completed = true
		return true
	}
	return false
}
