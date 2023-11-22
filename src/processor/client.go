package processor

import (
	"bufio"
	"dalun/commands"
	"net"
	"net/textproto"
	"time"
)

type Client struct {
	Connection    net.Conn
	ResultChannel chan commands.CommandResult
	CreatedAt     time.Time
}

func (c *Client) Start() {
	reader := bufio.NewReader(c.Connection)
	textReader := textproto.NewReader(reader)

	writer := bufio.NewWriter(c.Connection)
	textWriter := textproto.NewWriter(writer)
	defer c.Connection.Close()

	var errorCounter int = 0
	for {
		line, err := textReader.ReadLine()
		if err != nil {
			errorCounter++
			if errorCounter >= 5 {
				break
			}
			println("Error: " + err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		cmd := commands.Command{
			CMD:           line,
			ResultChannel: &c.ResultChannel,
		}

		processor.CommandChannel <- cmd
		result := <-c.ResultChannel
		if result.Success {
			textWriter.PrintfLine("", result.Data)
		} else {
			textWriter.PrintfLine("Error(%d):%s", result.Error.Code, result.Error.Message)
		}

	}
}
