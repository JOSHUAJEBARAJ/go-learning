## Errors

There are three type of errors

- Syntax error
- Runtime error
- semantic error

### Error Interface
- Error is an  type interface
- The method name, Error()
- The Error() method to return a string

On occuring the error we can do the following thing

- Return the error to the caller
- Log the error and continue execution
- Stop the execution of the program
- Ignore it (this is highly not recommended)
- Panic (only in very rare conditions, we will discuss this further later)

## Creating new error

```
errors.New("Somestring")
```

## Panic Working  

- The execution is stopped
- Any deferred functions in the panicking function will be called
- Any deferred functions in the stack of the panicking function will be called
- It will continue all the way up the stack until it reaches main()
- Statements after the panicking function will not execute
- The program then crashes

```
Defer inside the function
Defer outstide the function
panic: Panic
```

## Recover
- defer function should be used
- Ideally above the pani
## Errors Guidelines
 
