package models

import "time"

type Job struct {
	Id        int
	Data      []byte
	Delay     time.Duration
	CreatedAt time.Time
	ExpireAt  time.Time
}
