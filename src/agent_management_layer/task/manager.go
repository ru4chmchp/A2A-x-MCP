package task

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskManager struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func NewManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) CreateTask(taskType string, params map[string]interface{}, agentID string) *Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	id := uuid.New().String()
	now := time.Now()

	task := &Task{
		ID:        id,
		Type:      taskType,
		Params:    params,
		Status:    StatusCreated,
		CreatedAt: now,
		UpdatedAt: now,
		AgentID:   agentID,
	}

	tm.tasks[id] = task

	return task

}

func (tm *TaskManager) UpdateStatus(id string, status TaskStatus) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exits := tm.tasks[id]

	if !exits {
		return false
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	return true
}

func (tm *TaskManager) AddArtifact(id string, artifact Artifact) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exits := tm.tasks[id]

	if !exits {
		return false
	}

	task.Result = append(task.Result, artifact)
	task.UpdatedAt = time.Now()
	return true
}

func (tm *TaskManager) SetError(id string, errMsg string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	task, exits := tm.tasks[id]

	if !exits {
		return false
	}

	task.Status = StatusFailed
	task.Error = &errMsg
	task.UpdatedAt = time.Now()
	return true
}

func (tm *TaskManager) GetTaskByID(id string) (*Task, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exits := tm.tasks[id]
	return task, exits
}

func (tm *TaskManager) ListTasks() []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var result []*Task

	for _, task := range tm.tasks {
		result = append(result, task)
	}
	return result
}

// Lưu toàn bộ tasks từ TaskManager
func (tm *TaskManager) ExportTasks() []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var tasks []*Task
	for _, t := range tm.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

// Load lại task vào TaskManager
func (tm *TaskManager) ImportTasks(tasks []*Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for _, t := range tasks {
		tm.tasks[t.ID] = t
	}
}
