package task

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

type Task struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Payload    map[string]interface{} `json:"payload"`
	Attempts   int                    `json:"attempts"`
	MaxRetries int                    `json:"max_retries"`
	Status     Status                 `json:"status"`
	CreatedAt  time.Time              `json:"created_at"`
}

func New(taskType string, payload map[string]interface{}, maxRetries int) *Task {
	return &Task{
		ID:         uuid.New().String(),
		Type:       taskType,
		Payload:    payload,
		MaxRetries: maxRetries,
		Status:     StatusPending,
		Attempts:   0,
		CreatedAt:  time.Now(),
	}
}
