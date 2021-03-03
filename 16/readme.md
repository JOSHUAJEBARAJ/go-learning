https://www.youtube.com/watch?v=GJ5zvLZZyJM&list=PLv5GqfJa_Xn_DQa6lPL4GDfu2OBc0BVOc&index=21

## Concurrency

- Similar to thread
- forks the process
- main process should be alive

### Channel
- Pipe the data to the  go routines
- the read and write operation is blocked by default
- By default channels are bi directional

```
func main(){
	var c chan int // declaring the channel
	if c==nil{
		fmt.Println("Not intialized")
	}
	c=make(chan int ) // initailzing the value
}
```
## Deadlock
if there is an read there should be write and vice versa if there is only read there will be infinite wait
  
## Channel types
- unidirectional
- bi directional

**Sendonly**
- only can write

**Readonly**
- only can read but be aware this will lead to deadlock (there will be only read)

**Bi directional**
- can be converted into uni directional

### Closing and for loop

we can write the bulk data to the channel but while reading we have to iterate but the problem with iteration is we dont know when to stop 
- close will tell the sender when to notify the reader no more data (if we dont use it will create deadlock)
- using status flag we can find the sender notified or not (infinte loop)

### Buffer channel
- We can use buffer channel to perform the mutiple write 
- write will not be blocked until the buffer get fulled
- len use to find the current length of the buffer len(channel)
### Wait group
- add will create the waiter
- done will reduce the wait
- wait it will wait until wait=0
```
16.1 - go example
16.2 - creating channel
16.3 - channel block 
16.4 - channel 
16.5 - close
16.6 - length and buffered channel
16.7 - wait group
16.8 - race
16.9 mutex
```
