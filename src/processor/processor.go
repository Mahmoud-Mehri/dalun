package processor

import (
	"dalun/models"
	"time"
)

// var COMMANDS = []string{"use", "add", "delete", "get", ""}
type CommandResult struct {
	success bool
	data    *interface{}
	err     error
}

type Processor struct {
	DelayTicker    *time.Ticker
	QueueArray     []string
	Queues         map[string]*models.Queue
	CommandChannel chan string
}

func CheckDelays(q *models.Queue) {
	for id, job := range q.Delayed {
		jobTime := job.CreatedAt.Add(time.Second * job.Delay)
		if !jobTime.Before(time.Now()) {
			q.Ready[id] = job
			delete(q.Delayed, id)
		}
	}
}

func ProcessCommand(cmd string) *CommandResult {

	return nil
}

func StartProcessor() (*Processor, error) {
	p := Processor{}
	// Start Delay Checking Process
	p.DelayTicker = time.NewTicker(1 * time.Second)
	go func() {
		for range p.DelayTicker.C {
			for _, value := range p.QueueArray {
				q, found := p.Queues[value]
				if found {
					go CheckDelays(q)
				}
			}
		}
	}()

	return &p, nil
}
