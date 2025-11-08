package monitor

import (
	"encoding/json"
	"net/http"

	"worker-queue-system/internal/queue"
)

type Monitor struct {
	Queue *queue.Queue
	DLQ   *queue.Queue
}

func New(q *queue.Queue, dlq *queue.Queue) *Monitor {
	return &Monitor{
		Queue: q,
		DLQ:   dlq,
	}
}

func (m *Monitor) RegisterRoutes() {
	http.HandleFunc("/monitor/status", m.handleStatus)
}

func (m *Monitor) handleStatus(w http.ResponseWriter, _ *http.Request) {
	response := map[string]interface{}{
		"queue_length": m.Queue.Length(),
		"dlq_length":   m.DLQ.Length(),
		"queue_items":  m.Queue.Peek(),
		"dlq_items":    m.DLQ.Peek(),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
