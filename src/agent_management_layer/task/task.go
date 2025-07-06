package models

import (
	"sync"
	"time"
)

type Status string

const (
	StatusCreated    Status = "created"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type Task struct {
	TaskID    string                 `json:"task_id"`
	AgentID   string                 `json:"agent_id"`
	TaskType  string                 `json:"task_type"`
	Params    map[string]interface{} `json:"params"`
	Status    Status                 `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Result    interface{}            `json:"result,omitempty"`
	Error     *TaskError             `json:"error,omitempty"`
}

type TaskError struct {
	Message string `json:"message"`
	Trace   string `json:"trace,omitempty"`
}

type TaskManager struct {
	tasks map[string]*Task
	lock  sync.RWMutex
}

func newTaskManager()
