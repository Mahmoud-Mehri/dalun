package processor

import (
	"bufio"
	"dalun/models"
	"fmt"
	"net"
	"net/textproto"
	"time"
)

type Client struct {
	Connection    *net.Conn
	ResultChannel chan models.CommandResult
	CreatedAt     time.Time
	Running       bool
	Stopped       bool
}

func NewClient(con *net.Conn) (*Client, error) {
	client := Client{
		Connection:    con,
		ResultChannel: make(chan models.CommandResult),
		CreatedAt:     time.Now(),
		Running:       false,
		Stopped:       false,
	}

	return &client, nil
}

func (c *Client) Start() {
	reader := bufio.NewReader(*c.Connection)
	textReader := textproto.NewReader(reader)

	writer := bufio.NewWriter(*c.Connection)
	textWriter := textproto.NewWriter(writer)
	defer c.Close()

	c.Running = true
	var errorCounter int = 0
	for {
		if c.Stopped {
			break
		}

		line, err := textReader.ReadLine()
		if c.Stopped {
			break
		}

		if err != nil {
			errorCounter++
			if errorCounter >= 5 {
				break
			}
			println("Error: ", err.Error())

			if c.Stopped {
				break
			}

			time.Sleep(1 * time.Second)
			continue
		}

		println("New Command: ", line)

		cmd := models.Command{
			CMD:           line,
			ResultChannel: &c.ResultChannel,
		}

		fmt.Println("Command Created")
		if GlobalProcessor == nil {
			fmt.Println("Command Channel is Nil")
		}

		GlobalProcessor.CommandChannel <- cmd
		if c.Stopped {
			break
		}

		println("After Channel")

		result := <-c.ResultChannel
		if c.Stopped {
			break
		}

		if result.Success {
			textWriter.PrintfLine("Success - %v", result.Data)
		} else {
			textWriter.PrintfLine("Error(%d):%s", result.Error.Code, result.Error.Message)
		}

	}
}

func (c *Client) Stop() {
	if c.Running && !c.Stopped {
		c.Stopped = true
	}
}

func (c *Client) Close() {
	if c.Running {
		c.Stopped = true
		c.Running = false
	}
}
