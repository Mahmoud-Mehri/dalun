package models

type Queue struct {
	Id          int
	Name        string
	Delayed     map[int]Job
	Ready       map[int]Job
	Reserved    map[int]Job
	Ignored     map[int]Job
	Subscribers map[int]Subscriber
}
