package models

import "time"

type Job struct {
	Id        int
	Data      []byte
	Delay     time.Duration
	CreatedAt time.Time
	ExpireAt  time.Time
}

func NewJob(data []byte, delay int, expire time.Time) (*Job, error) {
	job := Job{}
	copy(job.Data, data)
	job.Delay = time.Duration(delay * 1000000)
	job.ExpireAt = expire

	return &job, nil
}
