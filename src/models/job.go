package models

import "time"

type Job struct {
	Id          uint
	Data        []byte
	CreatedAt   time.Time
	Delay       time.Duration
	ExpireAfter time.Duration
}

func NewJob(id uint, data []byte, delay int, expireAfter int) (*Job, error) {
	job := Job{
		Id: id,
	}
	copy(job.Data, data)
	job.Delay = time.Duration(int(time.Second) * delay)
	job.ExpireAfter = time.Duration(int(time.Second) * expireAfter)

	return &job, nil
}
