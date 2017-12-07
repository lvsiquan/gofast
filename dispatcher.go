package gofast

import (
	"os"
	"runtime"
	"strconv"
)

var MaxWorker = os.Getenv("MAX_WORKER")
var MaxQueue = os.Getenv("MAX_QUEUE")

// var JobQueue = make(chan Job)

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorker  int
	JobQueue   chan Job
	quit       chan bool
}

func NewDispatcher() *Dispatcher {
	maxWorker, err := strconv.Atoi(MaxWorker)
	if err != nil {
		maxWorker = runtime.NumCPU()
	}
	pool := make(chan chan Job, maxWorker)
	return &Dispatcher{WorkerPool: pool, MaxWorker: maxWorker, quit: make(chan bool), JobQueue: make(chan Job)}
}

func (d *Dispatcher) Run(handler JobHandler) {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool, handler, i)
		worker.start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func(job Job) {
				jobChan := <-d.WorkerPool
				jobChan <- job
			}(job)
		case <-d.quit:
			return
		}
	}
}

func (d *Dispatcher) Stop() {
	go func() {
		d.quit <- true
	}()
}
