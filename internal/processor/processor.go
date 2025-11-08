package processor

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"worker-queue-system/internal/task"
)

type Processor struct{}

func New() *Processor {
	return &Processor{}
}

func (p *Processor) Process(t *task.Task) error {
	fmt.Printf("[processor] Running task %s (%s)\n", t.ID, t.Type)

	switch t.Type {

	case "generate_report":
		return p.generateReport(t)

	case "compress_file":
		return p.compressFile(t)

	default:
		return errors.New("unknown task type: " + t.Type)
	}
}

func (p *Processor) generateReport(t *task.Task) error {
	name, ok := t.Payload["name"].(string)
	if !ok || name == "" {
		return errors.New("invalid payload: missing 'name'")
	}

	fmt.Printf("[report] Starting report generation: '%s'\n", name)

	fmt.Println("[report] Fetching data...")
	time.Sleep(400 * time.Millisecond)

	fmt.Println("[report] Processing statistics...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("[report] Rendering graphs...")
	time.Sleep(600 * time.Millisecond)

	fmt.Println("[report] Compiling final report...")
	time.Sleep(400 * time.Millisecond)

	fmt.Printf("[report] Report '%s' generated successfully (demo mode, no file created).\n", name)
	return nil
}

func (p *Processor) compressFile(t *task.Task) error {
	fileName, ok := t.Payload["file"].(string)
	if !ok || fileName == "" {
		return errors.New("invalid payload: missing 'file'")
	}

	fmt.Printf("[compress] Starting compression: '%s'\n", fileName)

	// Simulate file size
	fileSize := 20 + rand.Intn(80)
	fmt.Printf("[compress] Estimated size: %dMB\n", fileSize)

	totalBlocks := fileSize / 5
	for i := 1; i <= totalBlocks; i++ {
		fmt.Printf("[compress] Compressing block %d/%d...\n", i, totalBlocks)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("[compress] File '%s' compressed successfully (demo mode, no file created).\n", fileName)
	return nil
}
