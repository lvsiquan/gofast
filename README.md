# gofast
A worker pool written in Golang

# Usage
```
os.Setenv("MAX_WORKER", "20")
os.Setenv("MAX_QUEUE", "100")

d := NewDispatcher()
d.Run(func(job Job) { time.Sleep(1 * time.Second) })

for i := 0; i < 100; i++ {
	job := Job{Payload: "job" + strconv.Itoa(i)}
	d.JobQueue <- job
}
  
 ```
 
 # Benchmark
 > BenchmarkNewDispatcher-8   	  500000	      4499 ns/op	     557 B/op	       4 allocs/op
PASS
ok  	github.com/lvsiquan/gofast	2.412s
