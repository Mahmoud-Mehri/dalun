package models

import "time"

type Job struct {
	Id          int
	Data        []byte
	CreatedAt   time.Time
	Delay       time.Duration
	ExpireAfter time.Duration
}

func NewJob(data []byte, delay int, expireAfter int) (*Job, error) {
	job := Job{}
	copy(job.Data, data)
	job.Delay = time.Duration(int(time.Second) * delay)
	job.ExpireAfter = time.Duration(int(time.Second) * expireAfter)

	return &job, nil
}
