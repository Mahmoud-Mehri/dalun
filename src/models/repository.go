package models

import "errors"

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
func (repo *JobRepository) AddJob(qname string, job *Job) error {

	return nil
}
