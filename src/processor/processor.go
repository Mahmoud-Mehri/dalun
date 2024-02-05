package processor

import (
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
	Repository     *models.JobRepository
	CommandChannel chan models.Command
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

func ProcessCommand(cmd *models.Command) {
	cmdParts := strings.Split(cmd.CMD, " ")

	commandFunc := COMMAND_LIST[cmdParts[0]]
	var commandResult *models.CommandResult = nil
	if commandFunc != nil {
		commandResult = commandFunc(nil, cmd.CMD)
		cmd.ResultChannel <- *commandResult
	}
}

func StartProcessor(repo *models.JobRepository) (*Processor, error) {
	if processor != nil {
		return processor, nil
	}

	processor := Processor{}
	// Start Delay Checking Process
	processor.DelayTicker = time.NewTicker(DELAY_CHECK_TRESHOLD * time.Second)
	go func() {
		for range processor.DelayTicker.C {
			for _, value := range processor.Repository.QueueArray {
				q, found := processor.Repository.Queues[value]
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
			for _, value := range processor.Repository.QueueArray {
				q, found := processor.Repository.Queues[value]
				if found {
					go CheckExpires(q)
				}
			}
		}
	}()

	// Start Command Channel
	processor.CommandChannel = make(chan models.Command)
	go func() {
		for cmd := range processor.CommandChannel {
			go ProcessCommand(&cmd)
		}
	}()

	return &processor, nil
}
