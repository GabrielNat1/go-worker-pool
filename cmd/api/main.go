package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"worker-queue-system/cmd/worker"
	"worker-queue-system/internal/monitor"
	"worker-queue-system/internal/queue"
	"worker-queue-system/internal/task"
)

type TaskRequest struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

var (
	mainQueue  = queue.New()
	workerPool = worker.NewWorkerPool(3, mainQueue)
)

func main() {
	workerPool.Start()

	mon := monitor.New(mainQueue, workerPool.DLQ)
	mon.RegisterRoutes()

	http.HandleFunc("/task", handleCreateTask)

	fmt.Println("[api] Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		http.Error(w, "missing task type", http.StatusBadRequest)
		return
	}

	t := task.New(req.Type, req.Payload, 3)
	mainQueue.Push(t)

	response := map[string]string{
		"id":     t.ID,
		"status": "queued",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	fmt.Printf("[api] New task created: %s (%s)\n", t.ID, req.Type)
}
