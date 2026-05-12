package task


type TaskList struct {
	tasks map[int64]Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		tasks: make(map[int64]Task),
	}
}


func (tl *TaskList) Add(t Task) {
	tl.tasks[t.ID] = t
}


func (tl *TaskList) GetByID(id int64) (Task, bool) {
	t, exists := tl.tasks[id]
	return t, exists
}


func (tl *TaskList) GetAll() []Task {
	result := make([]Task, 0, len(tl.tasks))
	for _, t := range tl.tasks {
		result = append(result, t)
	}
	return result
}


func (tl *TaskList) Update(t Task) bool {
	if _, exists := tl.tasks[t.ID]; !exists {
		return false
	}
	tl.tasks[t.ID] = t
	return true
}


func (tl *TaskList) DeleteByID(id int64) bool {
	if _, exists := tl.tasks[id]; !exists {
		return false
	}
	delete(tl.tasks, id)
	return true
}


func (tl *TaskList) Count() int {
	return len(tl.tasks)
}
