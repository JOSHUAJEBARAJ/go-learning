## Concurrency

```
concurrent - two line one coffee machine
parallel - two line two coffee machine
```

Problem with the concurrent function is the main function doesn't wait until the go routines finish
## WaitGroup

by default each program runs one go routine called main goroutine think wait group as like the counter
wait.add will add the counter
wait.wait will tell to wait
wait.done will decrement the counter

### Race condition
this will happen when we use the share resources simultaneously
### Atomic package
### Mutex
### Channels
channel can be created with like a slice
```
ch := make(chan int)
```
ch <- 2 to assign the value
i:= <-ch to receive the value
```
16.1 concurrent-sum 0
16.2 - waitgroup
16.3 - atomic package
16.4 channels
16.5 two way channel
```