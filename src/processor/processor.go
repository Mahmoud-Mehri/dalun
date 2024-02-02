package processor

import (
	"dalun/commands"
	"dalun/models"
	"strings"
	"time"
)

const DELAY_CHECK_TRESHOLD = 1
const EXPIRE_CHECK_TRESHOLD = 1

var processor *Processor

type Processor struct {
	DelayTicker    *time.Ticker
	ExpireTicker   *time.Ticker
	QueueArray     []string
	Queues         map[string]*models.Queue
	CommandChannel chan commands.Command
}

// Checking Jobs in a Queue for reaching Delay time
func CheckDelays(q *models.Queue) {
	for id, job := range q.Delayed {
		readyTime := job.CreatedAt.Add(time.Second * job.Delay)
		if !readyTime.Before(time.Now()) {
			q.Ready[id] = job
			delete(q.Delayed, id)
		}
	}
}

// Checking Jobs in a Queue for Expiration
func CheckExpires(q *models.Queue) {
	for id, job := range q.Ready {
		if job.ExpireAfter == 0 {
			continue
		}

		expireTime := job.CreatedAt.Add(time.Second * job.ExpireAfter)
		if expireTime.Before(time.Now()) || expireTime.Equal(time.Now()) {
			delete(q.Ready, id)
		}
	}
}

func ProcessCommand(cmd *commands.Command) {
	cmdParts := strings.Split(cmd.CMD, " ")
	commandId := commands.COMMAND_LIST[cmdParts[0]]
	if commandId == 0 {
		println("Invalid command: %s", cmd.CMD)
	}

	commandFunc := commands.COMMAND_FUNCTIONS[commandId]
	var commandResult *commands.CommandResult = nil
	if commandFunc != nil {
		commandResult = commandFunc(cmd.CMD)
		cmd.ResultChannel <- *commandResult
	}
}

func StartProcessor() (*Processor, error) {
	if processor != nil {
		return processor, nil
	}

	processor := Processor{}
	// Start Delay Checking Process
	processor.DelayTicker = time.NewTicker(DELAY_CHECK_TRESHOLD * time.Second)
	go func() {
		for range processor.DelayTicker.C {
			for _, value := range processor.QueueArray {
				q, found := processor.Queues[value]
				if found {
					go CheckDelays(q)
				}
			}
		}
	}()

	// Start Expiration Checking Process
	processor.ExpireTicker = time.NewTicker(EXPIRE_CHECK_TRESHOLD * time.Second)
	go func() {
		for range processor.ExpireTicker.C {
			for _, value := range processor.QueueArray {
				q, found := processor.Queues[value]
				if found {
					go CheckExpires(q)
				}
			}
		}
	}()

	// Start Command Channel
	processor.CommandChannel = make(chan commands.Command)
	go func() {
		for cmd := range processor.CommandChannel {
			go ProcessCommand(&cmd)
		}
	}()

	return &processor, nil
}
