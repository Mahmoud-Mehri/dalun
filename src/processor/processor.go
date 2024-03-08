package processor

import (
	"dalun/models"
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

func ProcessCommand(repo *models.JobRepository, cmd *models.Command) {
	// fmt.Println("Processing Command")
	cmdParts := strings.Split(cmd.CMD, " ")

	commandFunc := COMMAND_LIST[cmdParts[0]]
	var commandResult *models.CommandResult = nil
	if commandFunc != nil {
		commandResult = commandFunc(repo, cmd.CMD)
		*cmd.ResultChannel <- *commandResult
	} else {
		commandResult = &models.CommandResult{
			Success: false,
			Data:    models.ERROR_INVALID_COMMAND_MSG,
			Error:   models.NewCommandError(models.ERROR_INVALID_COMMAND, models.ERROR_INVALID_COMMAND_MSG),
		}
		*cmd.ResultChannel <- *commandResult
		// println(models.ERROR_INVALID_COMMAND_MSG)
	}
}

func StartProcessor(repo *models.JobRepository) error {
	if GlobalProcessor != nil {
		return nil
	}

	GlobalProcessor = &Processor{
		Repository: repo,
	}

	// Start Delay Checking Process
	GlobalProcessor.DelayTicker = time.NewTicker(DELAY_CHECK_TRESHOLD * time.Second)
	go func() {
		for range GlobalProcessor.DelayTicker.C {
			var queueNames = GlobalProcessor.Repository.GetQueueNames()
			for _, value := range queueNames {
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
			var queueNames = GlobalProcessor.Repository.GetQueueNames()
			for _, value := range queueNames {
				q, found := GlobalProcessor.Repository.Queues[value]
				if found {
					go CheckExpires(q)
				}
			}
		}
	}()

	// Start Command Channel
	GlobalProcessor.CommandChannel = make(chan models.Command)
	go func() {
		for cmd := range GlobalProcessor.CommandChannel {
			println("Command Received on Processor:", cmd.CMD)
			go ProcessCommand(repo, &cmd)
		}
	}()

	return nil
}
