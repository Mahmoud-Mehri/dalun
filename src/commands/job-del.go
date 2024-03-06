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
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_COMMAND,
			Message: models.ERROR_INVALID_COMMAND_MSG,
		}

		return &result
	}

	if len(commandParts) < 3 {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}

		return &result
	}

	qname := commandParts[1]

	jobId, err := strconv.Atoi(commandParts[2])
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}

		return &result
	}

	err = repo.DeleteJob(qname, jobId)
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INTERNAL,
			Message: models.ERROR_INTERNAL_MSG,
		}

		return &result
	}

	result.Success = true
	result.Data = strconv.Itoa(jobId)

	return &result
}
