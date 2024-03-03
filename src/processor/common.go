package processor

import (
	"dalun/commands"
	"dalun/models"
)

var COMMAND_LIST = map[string]models.CommandFunction{
	"queue-add":  commands.AddQueueCommand,
	"queue-del":  commands.DelQueueCommand,
	"queue-info": nil,
	"queue-list": commands.ListQueueCommand,
	"job-add":    commands.AddJobCommand,
	"job-del":    commands.DelJobCommand,
	"job-get":    nil,
	"job-info":   nil,
	"job-list":   nil,
}
