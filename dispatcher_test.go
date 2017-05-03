package gofast

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func BenchmarkNewDispatcher(b *testing.B) {

	os.Setenv("MAX_WORKER", "20")
	os.Setenv("MAX_QUEUE", "100")

	d := NewDispatcher()
	d.Run(func(job Job) { time.Sleep(1 * time.Second) })

	for i := 0; i < b.N; i++ {
		job := Job{Payload: "job" + strconv.Itoa(i)}
		JobQueue <- job
	}

}
