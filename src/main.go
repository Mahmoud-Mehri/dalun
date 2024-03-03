package main

import (
	"dalun/models"
	"dalun/processor"
	"fmt"
	"net"
)

var clients []*processor.Client

func handleConnection(con *net.Conn) {
	client, err := processor.NewClient(con)
	if err != nil {
		fmt.Println("Error on creating new Client: ", err.Error())
		return
	}

	clients = append(clients, client)
	go client.Start()
}

func main() {
	const PORT = 8099
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", PORT))

	if err != nil {
		fmt.Printf("Error on listening port %d\n%s\n", PORT, err.Error())
		return
	}

	fmt.Println("Dalun is listening on port ", PORT)

	defer server.Close()

	repo := models.JobRepository{
		Queues: map[string]*models.Queue{},
	}

	err = processor.StartProcessor(&repo)
	if err != nil {
		fmt.Printf("Error on creating processor: %s", err.Error())
	}

	for {
		con, err := server.Accept()
		if err != nil {
			fmt.Println("Error on connection:", err.Error())
			continue
		}

		handleConnection(&con)
	}
}
