package commands

import (
	"dalun/models"
	"strings"
)

func AddQueueCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if !((len(commandParts) == 2) && (commandParts[0] == "queue-add")) {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: "Invalid Command Format",
		}
		return &result
	}

	queueName := commandParts[1]

	err := repo.AddQueue(queueName)
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INTERNAL,
			Message: err.Error(),
		}
	} else {
		result.Success = true
		// result.Data = "Queue added successfully"
	}

	return &result
}
