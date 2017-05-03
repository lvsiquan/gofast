package gofast

import (
	"strconv"
)

type JobHandler func(job Job)

type Job struct {
	Payload interface{}
}

type Worker struct {
	Name       string
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	handler    JobHandler
}

func NewWorker(workerpool chan chan Job, handler JobHandler, workerNo int) Worker {
	i_max_queue, err := strconv.Atoi(MaxQueue)
	if err != nil {
		i_max_queue = 10
	}
	return Worker{Name: "worker-" + strconv.Itoa(workerNo), WorkerPool: workerpool, JobChannel: make(chan Job, i_max_queue), quit: make(chan bool), handler: handler}
}

func (w Worker) start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// fmt.Printf("Worker is : %s,  ", w.Name)
				w.handler(job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}
