package models

import (
	"encoding/json"
)

type Task struct {
	Type   string          `json:"type" storm:"index"`
	Params json.RawMessage `json:"params,omitempty"`
}

type TaskRun struct {
	Task
	ID     string `storm:"id"`
	Status string
	Result RunResult
}
