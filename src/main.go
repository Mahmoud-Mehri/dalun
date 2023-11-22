package main

import (
	"dalun/commands"
	"dalun/processor"
	"fmt"
	"net"
	"time"
)

var clients []*processor.Client

func handleConnection(con net.Conn) {
	client := processor.Client{
		Connection:    con,
		ResultChannel: make(chan commands.CommandResult),
		CreatedAt:     time.Now(),
	}

	clients = append(clients, &client)
	go client.Start()
}

func main() {
	const PORT = 8099
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", PORT))

	if err != nil {
		fmt.Printf("Error on listening port %d\n%s\n", PORT, err.Error())
		return
	}

	defer server.Close()

	for {
		con, err := server.Accept()
		if err != nil {
			println("Error on connection:\n" + err.Error())
			continue
		}

		handleConnection(con)
	}
}
