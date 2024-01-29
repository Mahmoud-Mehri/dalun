package commands

type CommandFunction func(cmd string) *CommandResult

var COMMAND_LIST = map[string]int{
	"add-queue": 1,
	"del-queue": 2,
	"add-job":   3,
	"del-job":   4,
}
var COMMAND_FUNCTIONS = map[int]CommandFunction{
	1: AddJobCommand,
	2: DelJobCommand,
}

type CommandError struct {
	Code    int
	Message string
}

type Command struct {
	CMD           string
	ResultChannel chan CommandResult
}

type CommandResult struct {
	Success bool
	Data    *interface{}
	Error   CommandError
}
