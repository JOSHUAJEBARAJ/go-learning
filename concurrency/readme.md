## Go 

### sync.Waitgroup

A waitGroup waits for collection of go routines to finish

- [ ]  Create the waitgroup pointer `wg := &sync.WaitGroup{}`
- [ ]  Invoke the add method `wg.Add`
- [ ]  Invoke the done method 
- [ ]  Add the `wg.Wait` in the main

## Mutex

Mutal exclusion

```
go run --race mutex.go
```
- [ ] Create the mutex `m := &sync.Mutex{}`
- [ ] Add the mutex 
```
m.Lock()
x = x + 1
m.Unlock()
```

## Channel

Channel is used to pass the message between two go routines

So we don't have to worry about locking and unlocking 

Both go routines in independent of each other

### Overview
- Creating channel
- Buffered  channels
- Channel types
- Closing channel
- Control flows

## Creating channel

```
ch :=make(chan int)

// create buffeted channel
ch :=make(chan int,5)
```
Channels are **blocking** by default

```
	ch := make(chan int)

	fmt.Println(<-ch) // waiting the channel
	ch <- 2
```

## Channel type
- Bi directional
- Send only
- Recieve only

```

// recieve-only
func recieve(ch<- chan int){
    
}
// send only
func send(ch chan<- int){

}
```

## closing the channel

- Close via the built-in

```
close(ch)
```

- Closing in the sending end will create the deadlock
- *Closing on the recieving end doesn't create deadlock*

> sending message in closed channel panic recieving message in the closed channel will get zero value

## Control flow
- If statements
- For loops
- Select statement


```
msg, ok := <-ch
		if ok {
			fmt.Println(msg, ok)
		}
```

## For statement
Sometimes the sending side doesnt know how many message it will be generating
```
for msg := range ch {
			fmt.Println(msg)
		}
```
> Note we have to close at the sending side

## Select statement

```
ch1 := make (chan int)
ch2 := make( chan string)

select{
    case i:= <-ch:
    case ch <- "ch":

    default:
    // non blocking select 
}
```