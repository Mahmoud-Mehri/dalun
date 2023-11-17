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
	ExpireTicker   *time.Ticker
	QueueArray     []string
	Queues         map[string]*models.Queue
	CommandChannel chan string
}

func CheckDelays(q *models.Queue) {
	for id, job := range q.Delayed {
		readyTime := job.CreatedAt.Add(time.Second * job.Delay)
		if !readyTime.Before(time.Now()) {
			q.Ready[id] = job
			delete(q.Delayed, id)
		}
	}
}

func CheckExpires(q *models.Queue) {
	for id, job := range q.Ready {
		if job.ExpireAfter == 0 {
			continue
		}

		expireTime := job.CreatedAt.Add(time.Second * job.ExpireAfter)
		if expireTime.Before(time.Now()) {
			delete(q.Ready, id)
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

	p.ExpireTicker = time.NewTicker(1 * time.Second)
	go func() {
		for range p.ExpireTicker.C {
			for _, value := range p.QueueArray {
				q, found := p.Queues[value]
				if found {
					go CheckExpires(q)
				}
			}
		}
	}()

	return &p, nil
}
