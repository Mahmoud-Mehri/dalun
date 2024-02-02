package commands

import "strings"

func AddJobCommand(cmd string) *CommandResult {
	result := CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] != "add-job" {
		result.Success = false
	}

	return &result
}
