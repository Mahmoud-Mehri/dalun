package models

import "net"

type Subscriber struct {
	Id int
	Connection *net.Conn
}