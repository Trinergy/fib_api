# fib_api
This app has 3 endpoints which allow you to step through the fibonacci sequence at your own pace 🌙🚶‍♂️🔢

## Endpoints
* `/current` returns the value of current sequence
* `/next` returns the value of next sequence and increments the index
* `/previous` returns the values of the previous sequence and decrements the index

```
Example:
current -> 0
next -> 1
next -> 1
next -> 2
previous -> 1
```

## How to use
```
go run main.go
// go to localhost:8080
```

## Benchmark
Benchmarked using [Docker](https://www.docker.com) and [wrk](https://github.com/wg/wrk) based on the following specifications:
* 512MB RAM
* 1 CPU

### Results
```
Running 5s test @ http://127.0.0.1:8080/current
  1 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    30.74ms   25.42ms 119.35ms   68.97%
    Req/Sec     3.69k   663.34     4.96k    70.00%
  18392 requests in 5.02s, 3.44MB read
Requests/sec:   3664.52
Transfer/sec:    701.41KB

Running 5s test @ http://127.0.0.1:8080/previous
  1 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    41.00ms   30.35ms 182.62ms   61.27%
    Req/Sec     2.47k   535.60     3.55k    66.00%
  12326 requests in 5.03s, 1.41MB read
Requests/sec:   2452.75
Transfer/sec:    287.75KB

Running 5s test @ http://127.0.0.1:8080/next
  1 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    42.63ms   34.17ms 210.69ms   65.07%
    Req/Sec     2.46k   367.12     3.42k    74.00%
  12235 requests in 5.02s, 1.96MB read
  Non-2xx or 3xx responses: 6247
Requests/sec:   2439.63
Transfer/sec:    399.50KB
```
Note: `Non-2xx or 3xx responses` from `/next` endpoint are expected for handling integer overflow
### How to benchmark
```
// Setup
docker build -t fib-api .
docker run --mount type=tmpfs,destination=/tmp/ -p 8080:8080 -d --cpus=1 --memory="512m" fib-api

// Benchmark
wrk -t1 -c100 -d5s http://127.0.0.1:8080/current
```
### Comments
* Using [endless](https://github.com/fvbock/endless) to handle graceful restarts, but would probably use a real supervisor in production
* Errors in HTTP handlers print to standard error, but would probably be sent to a monitoring system in production
