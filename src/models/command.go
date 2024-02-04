package models

type CommandFunction func(cmd string) *CommandResult

type Command struct {
	CMD           string
	ResultChannel chan CommandResult
}

type CommandResult struct {
	Success bool
	Data    *interface{}
	Error   CommandError
}
