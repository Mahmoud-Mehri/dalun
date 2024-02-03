package commands

import (
	"dalun/models"
	"strings"
)

func AddQueueCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if !((len(commandParts) == 2) && (commandParts[0] == "add-queue")) {
		result.Success = false
		result.Error = models.CommandError{
			Code:    1,
			Message: "Invalid Command",
		}
		return &result
	}

	queueName := commandParts[1]

	repo.AddQueue(queueName)

	return &result
}
