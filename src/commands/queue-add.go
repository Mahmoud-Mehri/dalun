package commands

import (
	"dalun/models"
	"strings"
)

func AddQueueCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] != "queue-add" {
		result.Success = false
		result.Error = models.NewCommandError(models.ERROR_INVALID_COMMAND, models.ERROR_INVALID_COMMAND_MSG)

		return &result
	}

	if len(commandParts) != 2 {
		result.Success = false
		result.Error = models.NewCommandError(models.ERROR_INVALID_FORMAT, models.ERROR_INVALID_FORMAT_MSG)

		return &result
	}

	queueName := commandParts[1]

	err := repo.AddQueue(queueName)
	if err != nil {
		result.Success = false
		result.Error = err
	} else {
		result.Success = true
		result.Data = queueName
	}

	return &result
}
