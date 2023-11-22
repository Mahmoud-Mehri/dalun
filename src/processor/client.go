package processor

import (
	"dalun/commands"
	"net"
	"time"
)

type Client struct {
	Connection     net.Conn
	CommandChannel chan string
	ResultChannel  chan commands.CommandResult
	CreatedAt      time.Time
}
