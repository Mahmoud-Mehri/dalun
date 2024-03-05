package commands

import (
	"dalun/models"
	"strings"
)

func DelQueueCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if !((len(commandParts) == 2) && (commandParts[0] == "queue-del")) {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}
		return &result
	}

	queueName := commandParts[1]

	err := repo.DeleteQueue(queueName)
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INTERNAL,
			Message: models.ERROR_INTERNAL_MSG,
		}
	} else {
		result.Success = true
		result.Data = "QUEUE_DELETED"
	}

	return &result
}
