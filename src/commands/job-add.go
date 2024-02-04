package commands

import (
	"dalun/models"
	"strings"
)

func AddJobCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] != "job-add" {
		result.Success = false
	}

	return &result
}
