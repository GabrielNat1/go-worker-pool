package worker

import (
	"fmt"
	"time"

	"worker-queue-system/internal/processor"
	"worker-queue-system/internal/queue"
)

type Pool struct {
	Queue     *queue.Queue
	DLQ       *queue.Queue
	Processor *processor.Processor
	Workers   int
}

func NewWorkerPool(workers int, q *queue.Queue) *Pool {
	return &Pool{
		Queue:     q,
		DLQ:       queue.New(),
		Processor: processor.New(),
		Workers:   workers,
	}
}

func (wp *Pool) Start() {
	fmt.Printf("[worker-pool] Starting %d workers...\n", wp.Workers)

	for i := 0; i < wp.Workers; i++ {
		go wp.workerLoop(i + 1)
	}
}

func (wp *Pool) workerLoop(id int) {
	fmt.Printf("[worker-%d] started\n", id)
	for {
		t, err := wp.Queue.Pop()
		if err != nil {
			// queue empties → small wait
			time.Sleep(300 * time.Millisecond)
			continue
		}

		fmt.Printf("[worker-%d] processing task %s (%s)\n", id, t.ID, t.Type)

		err = wp.Processor.Process(t)
		if err != nil {
			fmt.Printf("[worker-%d] error: %s\n", id, err.Error())
			t.Attempts++

			if t.Attempts >= t.MaxRetries {
				fmt.Printf("[worker-%d] moving to DLQ: %s\n", id, t.ID)
				wp.DLQ.Push(t)
			} else {
				fmt.Printf("[worker-%d] requesting task %s (attempt %d)\n", id, t.ID, t.Attempts)
				wp.Queue.Push(t)
			}

			continue
		}

		fmt.Printf("[worker-%d] task %s completed successfully\n", id, t.ID)
	}
}
