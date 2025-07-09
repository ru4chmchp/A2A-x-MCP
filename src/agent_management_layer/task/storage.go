package task

import (
	"encoding/json"
	"os"
	"sync"
)

type TaskStorage struct {
	mu       sync.RWMutex
	FilePath string
}

func (ts *TaskStorage) Load() ([]*Task, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	file, err := os.Open(ts.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var tasks []*Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskStorage) Save(tasks []*Task) error {
	ts.mu.RLock()
	defer ts.mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}
	err = os.WriteFile(ts.FilePath, data, 0664)
	return err
}
