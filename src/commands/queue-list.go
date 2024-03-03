package commands

import (
	"dalun/models"
	"strings"
)

func ListQueueCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	// commandParts := strings.Split(cmd, " ")
	if cmd != "queue-list" {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_COMMAND,
			Message: "Invalid Command",
		}
	}

	result.Success = true
	result.Data = strings.Join(repo.QueueArray, "\n")

	return &result
}
