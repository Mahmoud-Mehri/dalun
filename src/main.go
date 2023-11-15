package main

import (
	"fmt"
	"net"
)

func handleConnection(con net.Conn) {
	for {
		buf := make([]byte, 1024)
		len, err := con.Read(buf)
		if err != nil {
			println("Error on reading command")
			break
		}

	}

	err := con.Write([]byte("ERROR"))
	if err != nil {
		println("Error on writing ERROR command")
	}
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

		go handleConnection(con)
	}
}
