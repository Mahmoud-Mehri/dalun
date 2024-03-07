package models

import (
	"errors"
)

type JobRepository struct {
	queueArray []string
	Queues     map[string]*Queue
}

// Get Queue List
func (repo *JobRepository) GetQueueNames() []string {
	return repo.queueArray
}

// Adding new Queue
func (repo *JobRepository) AddQueue(qname string) error {
	q, err := NewQueue(qname)
	if err != nil {
		return err
	}

	println("Queue Created:", qname)

	repo.Queues[qname] = q
	repo.queueArray = append(repo.queueArray, qname)

	println("Queue Added:", qname)

	return nil
}

func (repo *JobRepository) DeleteQueue(qname string) error {
	if repo.Queues[qname] != nil {
		delete(repo.Queues, qname)
		for i, val := range repo.queueArray {
			if val == qname {
				repo.queueArray = append(repo.queueArray[:i], repo.queueArray[i+1:]...)
				break
			}
		}

		return nil
	}

	return errors.New(ERROR_QUEUE_NOTFOUND_MSG)
}

// Adding new Job
func (repo *JobRepository) AddJob(qname string, data []byte, delay int, expire int) (int, error) {
	if repo.Queues[qname] == nil {
		return 0, errors.New(ERROR_QUEUE_NOTFOUND_MSG)
	}

	job, err := repo.Queues[qname].AddNewJob(data, delay, expire)
	if err != nil {
		return 0, err
	}

	return job.Id, nil
}

// Delete new Job
func (repo *JobRepository) DeleteJob(qname string, jobId int) error {

	return nil
}
