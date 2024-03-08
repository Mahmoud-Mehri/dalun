package models

import (
	"sync"
)

type Queue struct {
	Id        uint
	Name      string
	LastJobId uint

	DelayLocker sync.Mutex
	Delayed     map[uint]*Job

	ReadyLocker sync.Mutex
	Ready       map[uint]*Job

	ReservedLocker sync.Mutex
	Reserved       map[uint]*Job

	IgnoredLocker sync.Mutex
	Ignored       map[uint]*Job

	SubscriberLocker sync.Mutex
	Subscribers      map[uint]*Subscriber
}

func NewQueue(name string) *Queue {
	queue := Queue{
		LastJobId:   0,
		Delayed:     map[uint]*Job{},
		Ready:       map[uint]*Job{},
		Reserved:    map[uint]*Job{},
		Ignored:     map[uint]*Job{},
		Subscribers: map[uint]*Subscriber{},
	}
	queue.Name = name

	return &queue
}

func (q *Queue) AddNewJob(data []byte, delay int, expireAfter int) (Job *Job, Err *CommandError) {
	defer func() {
		if err := recover(); err != nil {
			Err = NewCommandError(ERROR_INTERNAL, ERROR_INTERNAL_MSG)
		}
	}()

	newJobId := q.LastJobId + 1
	job := NewJob(newJobId, data, delay, expireAfter)

	if job.Delay > 0 {
		q.DelayLocker.Lock()
		q.Delayed[job.Id] = job
		q.DelayLocker.Unlock()
	} else {
		q.ReadyLocker.Lock()
		q.Ready[job.Id] = job
		q.ReadyLocker.Unlock()
	}

	return job, nil
}

func (q *Queue) DeleteJob(id uint) (Err *CommandError) {
	defer func() {
		if err := recover(); err != nil {
			Err = NewCommandError(ERROR_INTERNAL, ERROR_INTERNAL_MSG)
		}
	}()

	if q.Ready[id] == nil {
		if q.Delayed[id] == nil {
			return NewCommandError(ERROR_JOB_NOTFOUND, ERROR_JOB_NOTFOUND_MSG)
		} else {
			q.Delayed[id] = nil
			return nil
		}
	}

	q.Ready[id] = nil
	return nil
}
