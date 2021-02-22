## Basic Debugging

### Printf

```
% t - true or false
% d - decimal
% x - hex 
% b - binary 
```

## Additional Options for Formatting

```
%.2f - 2 decimal
% .3f - 3 decimal point
```

The number before the d b x specify the space
```
fmt.Printf("%10.d %15.b %10.x\n",i,i,i) 
```
## Logging

- uses "log" packages
- same as fmt

```
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
```
Date  time filename

### Log Fatal Errors

fatal exit with non 0 

**Fatal vs panic**

Fatal exits with non 0 while panic has the defer option