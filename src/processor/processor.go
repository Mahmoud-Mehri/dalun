package processor

import (
	"dalun/models"
	"fmt"
	"strings"
	"time"
)

const DELAY_CHECK_TRESHOLD = 1
const EXPIRE_CHECK_TRESHOLD = 1

var GlobalProcessor *Processor

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
	fmt.Println("Processing Command")
	cmdParts := strings.Split(cmd.CMD, " ")

	commandFunc := COMMAND_LIST[cmdParts[0]]
	var commandResult *models.CommandResult = nil
	if commandFunc != nil {
		commandResult = commandFunc(nil, cmd.CMD)
		*cmd.ResultChannel <- *commandResult
	}
}

func StartProcessor(repo *models.JobRepository) error {
	if GlobalProcessor != nil {
		return nil
	}

	GlobalProcessor = &Processor{
		Repository: repo,
	}

	fmt.Println("Processor Created")

	// Start Delay Checking Process
	GlobalProcessor.DelayTicker = time.NewTicker(DELAY_CHECK_TRESHOLD * time.Second)
	go func() {
		for range GlobalProcessor.DelayTicker.C {
			for _, value := range GlobalProcessor.Repository.QueueArray {
				q, found := GlobalProcessor.Repository.Queues[value]
				if found {
					go CheckDelays(q)
				}
			}
		}
	}()

	// Start Expiration Checking Process
	GlobalProcessor.ExpireTicker = time.NewTicker(EXPIRE_CHECK_TRESHOLD * time.Second)
	go func() {
		for range GlobalProcessor.ExpireTicker.C {
			for _, value := range GlobalProcessor.Repository.QueueArray {
				q, found := GlobalProcessor.Repository.Queues[value]
				if found {
					go CheckExpires(q)
				}
			}
		}
	}()

	fmt.Println("Before Command Channel")

	// Start Command Channel
	GlobalProcessor.CommandChannel = make(chan models.Command)
	go func() {
		fmt.Println("Processor Channel Function")
		for cmd := range GlobalProcessor.CommandChannel {
			println("Command Received on Processor")
			go ProcessCommand(&cmd)
		}
	}()

	return nil
}
