## Function

Syntax

```
    func name (variable type) type 
```

> Use camelcase

## Difference between argument and the parameter

When we define function is called parameter when we use it called argument

## Naked Returns

Second value returned from the function is mostly the error

## Variadic Function

- Multiple parameter is called **pack operator**
- must be in last place
- each function has one vardiac argument 

```
parameter --> argument --> slice
```

> if we want to pass the slice then we have to use the ... operator to the slice

## Anonymous Functions

### Method-1
```
func (){
    fmt.println("Hello")
}
```
### Method-2
```
m:="Joshua"
func(str string){
    fmt.Println(str)
}(m)
```

### Method-3

```
f:=func(){
    fmt.Println("variable")
}
f()
```

### Closure

- Some times we want to change the variable value only by certain function then we use the closure with anonomouyus
- The return type is anonmyous function
- closure function will have no parameter
- 2 return statement will be present

## Function Types 

To-study

```
type name func()
```

## defer
- will be done at the end
- uses LIFO
- Used to database close