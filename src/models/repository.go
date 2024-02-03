package models

type JobRepository struct {
	QueueArray []string
	Queues     map[string]*Queue
}

// Adding new Queue
func (p *JobRepository) AddQueue(qname string) {
	q := Queue{}
	p.Queues[qname] = &q
}

// Adding new Job
func (p *JobRepository) AddJob(qname string, job *Job) {

}
