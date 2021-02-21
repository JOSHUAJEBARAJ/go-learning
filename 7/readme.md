## Interface

```
type <name> Interface{
method set()
}
```
> we can have multiple methods

### Advantages of Implementing Interfaces Implicitly

There are two advantages of Interface in Go with no implicit 

- When there is method change in the interface we have to remove the implicit keyword everywhere 
- You can implement the interface in the other languages
  
  > eg stringer
```
  type Stringer interface {

  String() string

}
```

```
interface in the other packages  ---> struct <-- interface in the current package
```
### Duck typing - 6 steps

1. Declare the struct
2. Declare the interface
3. Create the struct variable
4. Implement the interface
5. Create the function with the Interface 
6. Call the interface method with the function

### poly morphism

1. Declare the different struct
2. Declare the interface
3. Create the struct variable for different struct
4. Implement the interface for different struct
5. Create the different function with the struct Interface 
6. Call the interface method with the function

> Use the vardiac function to call the interface

## Empty Interface

*every type implement the empty interface*

Used in the below
- Type switching
- Type assertion
- Packages

For example refer 4.20 and 4.21