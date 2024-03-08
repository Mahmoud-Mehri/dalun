package commands

import (
	"dalun/models"
	"strconv"
	"strings"
)

func AddJobCommand(repo *models.JobRepository, cmd string) *models.CommandResult {
	result := models.CommandResult{}

	commandParts := strings.Split(cmd, " ")
	if commandParts[0] != "job-add" {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_COMMAND,
			Message: models.ERROR_INVALID_COMMAND_MSG,
		}

		return &result
	}

	if len(commandParts) < 5 {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}

		return &result
	}

	qname := commandParts[1]

	delay, err := strconv.Atoi(commandParts[2])
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}

		return &result
	}

	expire, err := strconv.Atoi(commandParts[3])
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INVALID_FORMAT,
			Message: models.ERROR_INVALID_FORMAT_MSG,
		}

		return &result
	}

	var data []byte = []byte(strings.Join(commandParts[4:], " "))
	jobId, err := repo.AddJob(qname, data, delay, expire)
	if err != nil {
		result.Success = false
		result.Error = models.CommandError{
			Code:    models.ERROR_INTERNAL,
			Message: models.ERROR_INTERNAL_MSG,
		}

		return &result
	}

	result.Success = true
	result.Data = strconv.FormatUint(uint64(jobId), 10)

	return &result
}
