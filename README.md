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
Benchmarked with [Docker](https://www.docker.com) and [wrk](https://github.com/wg/wrk) according to the following specifications:
* 512MB RAM
* 1 CPU

### Results
```
Running 5s test @ http://127.0.0.1:8080/current
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.22ms   11.71ms  66.62ms   83.26%
    Req/Sec     3.03k   398.01     3.99k    82.00%
  15117 requests in 5.02s, 2.83MB read
Requests/sec:   3013.67
Transfer/sec:    576.84KB

Running 5s test @ http://127.0.0.1:8080/previous
  1 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.74ms   13.22ms  57.76ms   82.41%
    Req/Sec     2.42k   276.02     3.19k    76.47%
  12297 requests in 5.10s, 1.37MB read
Requests/sec:   2411.15
Transfer/sec:    275.95KB

Running 5s test @ http://127.0.0.1:8080/next
  1 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.79ms   16.11ms  76.43ms   81.36%
    Req/Sec     2.59k   314.92     3.48k    62.75%
  13127 requests in 5.10s, 2.51MB read
  Non-2xx or 3xx responses: 12229
Requests/sec:   2573.94
Transfer/sec:    504.25KB
```
Note: `Non-2xx or 3xx responses` from `/next` are expected responses for handling integer overflow
### How to benchmark
```
docker build -t fib-api .
docker run --mount type=tmpfs,destination=/tmp/ -p 8080:8080 -d --cpus=1 --memory="512m" fib-api
wrk -t1 -c10 -d5s http://127.0.0.1:8080/current
```
### Comments
Used [endless](https://github.com/fvbock/endless) to handle gracefully restarts, but would probably use a real supervisor in production
