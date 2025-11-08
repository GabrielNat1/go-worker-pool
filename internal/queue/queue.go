package queue

import (
	"errors"
	"sync"
	"worker-queue-system/internal/task"
)

type Queue struct {
	mu    sync.Mutex
	tasks []*task.Task
}

func New() *Queue {
	return &Queue{
		tasks: []*task.Task{},
	}
}

func (q *Queue) Push(t *task.Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(q.tasks, t)
}

func (q *Queue) Pop() (*task.Task, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return nil, errors.New("queue is empty")
	}

	t := q.tasks[0]
	q.tasks = q.tasks[1:]

	return t, nil
}

func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.tasks)
}

func (q *Queue) Peek() []*task.Task {
	q.mu.Lock()
	defer q.mu.Unlock()

	copied := make([]*task.Task, len(q.tasks))
	copy(copied, q.tasks)
	return copied
}
