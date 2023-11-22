package commands

type CommandFunction func(cmd *Command) *CommandResult

var COMMAND_LIST = [...]string{"add", "delete"}
var COMMAND_FUNCTIONS = map[string]CommandFunction{
	"add":    AddJobCommand,
	"delete": DeleteJobCommand,
}

type CommandError struct {
	Code    int
	Message string
}

type Command struct {
	CMD           string
	ResultChannel *chan CommandResult
}

type CommandResult struct {
	Success bool
	Data    *interface{}
	Error   CommandError
}
