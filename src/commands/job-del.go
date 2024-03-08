package commands

import (
	"dalun/models"
	"strconv"
	"strings"
)

func DelJobCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] == "job-del" {
		result.Success = false
		result.Error = models.NewCommandError(models.ERROR_INVALID_COMMAND, models.ERROR_INVALID_COMMAND_MSG)

		return &result
	}

	if len(commandParts) < 3 {
		result.Success = false
		result.Error = models.NewCommandError(models.ERROR_INVALID_FORMAT, models.ERROR_INVALID_FORMAT_MSG)

		return &result
	}

	qname := commandParts[1]

	jobId, err := strconv.ParseUint(commandParts[2], 10, 32)
	if err != nil {
		result.Success = false
		result.Error = models.NewCommandError(models.ERROR_INVALID_FORMAT, models.ERROR_INVALID_FORMAT_MSG)

		return &result
	}

	e := repo.DeleteJob(qname, uint(jobId))
	if e != nil {
		result.Success = false
		result.Error = e
		return &result
	}

	result.Success = true
	result.Data = strconv.FormatUint(jobId, 10)

	return &result
}
