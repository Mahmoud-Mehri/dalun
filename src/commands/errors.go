package commands

type CommandError struct {
	Code    int
	Message string
}

const (
	ERROR_INVALID_COMMAND = 1
	ERROR_INVALID_FORMAT  = 2
)
