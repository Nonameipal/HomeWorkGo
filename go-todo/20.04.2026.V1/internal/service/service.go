package service

import (
	"Todo/internal/models"
	"Todo/internal/repository"
)


type TaskService struct {
	repo *repository.TaskRepository
}


func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}


func (s *TaskService) GetAllTasks() []*models.Task {
	return s.repo.GetAll()
}

func (s *TaskService) AddTask(description string) int {
	return s.repo.Add(description)
}

func (s *TaskService) DeleteTask(id int) bool {
	return s.repo.Delete(id)
}

func (s *TaskService) CompleteTask(id int) bool {
	return s.repo.Complete(id)
}
