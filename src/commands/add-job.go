package commands

import (
	"dalun/models"
	"strings"
)

func AddJobCommand(cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] != "add-job" {
		result.Success = false
	}

	return &result
}
