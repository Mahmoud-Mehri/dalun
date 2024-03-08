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
		result.Error = models.NewCommandError(models.ERROR_INVALID_COMMAND, models.ERROR_INVALID_COMMAND_MSG)
		return &result
	}

	result.Success = true
	result.Data = strings.Join(repo.GetQueueNames(), "\n")

	return &result
}
