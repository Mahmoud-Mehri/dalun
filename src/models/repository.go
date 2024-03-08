package models

type JobRepository struct {
	queueArray []string
	Queues     map[string]*Queue
}

// Get Queue List
func (repo *JobRepository) GetQueueNames() []string {
	return repo.queueArray
}

// Adding new Queue
func (repo *JobRepository) AddQueue(qname string) *CommandError {
	if repo.Queues[qname] != nil {
		return NewCommandError(ERROR_QUEUE_DUPLICATE, ERROR_QUEUE_DUPLICATE_MSG)
	}

	q := NewQueue(qname)

	println("Queue Created:", qname)

	repo.Queues[qname] = q
	repo.queueArray = append(repo.queueArray, qname)

	println("Queue Added:", qname)

	return nil
}

func (repo *JobRepository) DeleteQueue(qname string) *CommandError {
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

	return NewCommandError(ERROR_QUEUE_NOTFOUND, ERROR_QUEUE_NOTFOUND_MSG)
}

// Adding new Job
func (repo *JobRepository) AddJob(qname string, data []byte, delay int, expire int) (uint, *CommandError) {
	if repo.Queues[qname] == nil {
		return 0, NewCommandError(ERROR_QUEUE_NOTFOUND, ERROR_QUEUE_NOTFOUND_MSG)
	}

	job, err := repo.Queues[qname].AddNewJob(data, delay, expire)
	if err != nil {
		return 0, err
	}

	return job.Id, nil
}

// Delete Job
func (repo *JobRepository) DeleteJob(qname string, jobId uint) *CommandError {
	if repo.Queues[qname] == nil {
		return NewCommandError(ERROR_QUEUE_NOTFOUND, ERROR_QUEUE_NOTFOUND_MSG)
	}

	err := repo.Queues[qname].DeleteJob(jobId)
	if err != nil {
		return err
	}

	return nil
}
