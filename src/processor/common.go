package processor

import (
	"dalun/commands"
	"dalun/models"
)

var COMMAND_LIST = map[string]int{
	"queue-add": 1,
	"queue-del": 2,
	"job-add":   3,
	"job-del":   4,
}
var COMMAND_FUNCTIONS = map[int]models.CommandFunction{
	1: commands.AddJobCommand,
	2: commands.DelJobCommand,
}
