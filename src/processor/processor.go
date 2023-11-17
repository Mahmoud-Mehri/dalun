package processor

import (
	"dalun/models"
	"time"
)

// var COMMANDS = []string{"use", "add", "delete", "get", ""}

type Processor struct {
	DelayTicker *time.Ticker
	QueueArray  []string
	Queues      map[string]*models.Queue
}

func CheckDelays(q *models.Queue) {
	if len(q.Delayed) > 0 {

	}
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
