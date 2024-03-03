package models

type JobRepository struct {
	QueueArray []string
	Queues     map[string]*Queue
}

// Adding new Queue
func (repo *JobRepository) AddQueue(qname string) error {
	q, err := NewQueue(qname)
	if err != nil {
		return err
	}

	repo.Queues[qname] = q
	repo.QueueArray = append(repo.QueueArray, qname)

	return nil
}

// Adding new Job
func (repo *JobRepository) AddJob(qname string, job *Job) error {

	return nil
}
