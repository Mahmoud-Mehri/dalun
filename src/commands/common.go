package commands

var COMMAND_LIST = [...]string{"add", "delete"}

type Command struct {
	CMD           string
	ResultChannel *chan CommandResult
}

type CommandResult struct {
	Success bool
	Data    *interface{}
	Error   error
}
