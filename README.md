# gofast
A worker pool written by Golang

# Usage
```
os.Setenv("MAX_WORKER", "20")
os.Setenv("MAX_QUEUE", "100")

d := NewDispatcher()
d.Run(func(job Job) { time.Sleep(1 * time.Second) })

for i := 0; i < 100; i++ {
	job := Job{Payload: "job" + strconv.Itoa(i)}
	JobQueue <- job
}
  
 ```
 
 # Benchmark
 > BenchmarkNewDispatcher-4          300000              7291 ns/op             622 B/op          5 allocs/op 
PASS
ok      github.com/gofast       2.878s
