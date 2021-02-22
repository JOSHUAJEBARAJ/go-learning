## Time


**Sleep function**
```
time.Sleep(2 * time.Second)
```


**Finding day and Time**
```
Day := time.Now().Weekday()
Hour := time.Now().Hour()
```
### Comparing Time

There are three functions
today before tommorow
tommorow After today
*tbt*

## Duration Calculation

- hours 
- minutes
- seconds
- nanoseconds


```
duration:=end.Sub(start)
```

## Formatting Time
- using parse
- using format
  
### Using Parse
- RFC3339
- UnixDate
- ANSIC

> LoadLocation() to load the time zone In() function of the time package to, say, convert that value to a specific time zone's value.

```
current time in
```
```
eg-1 Simple time
10.1 - function return the time stamp
eg-2 comparison
10.2 - Duration
eg-3 parse example
10.3 Time calculators
```