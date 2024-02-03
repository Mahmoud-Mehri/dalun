package processor

import (
	"dalun/commands"
	"dalun/models"
)

var COMMAND_LIST = map[string]int{
	"add-queue": 1,
	"del-queue": 2,
	"add-job":   3,
	"del-job":   4,
}
var COMMAND_FUNCTIONS = map[int]models.CommandFunction{
	1: commands.AddJobCommand,
	2: commands.DelJobCommand,
}
