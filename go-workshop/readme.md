# Go workshop

## Keywords

**Dependencies** - import package

**condition** - if else switch case fallthrough break default goto 

**loop** - for range continue

**Type** - var func interface struct chan const type map make

**Misc** - defer go return panic recover  


## Environment

**build** - compiles the binary

**fmt** - formats

**test** - runs the test

**tool** - launchers

**get** download dependencies

**mod** manages dependencies

### Running and Building the program

```
GOOS=linux GOARCH=arm go build -o toto
```

Go support multiple import

## Export an identifier

There is no private and public We can use the Upper case to denote public 

> Global variables can be used or not

## Function
- may single or multiple value
- {  has to be in the same line

## Standard variable declaration

```
var name type = value
```

> We can also short hand but inside the function only

We can also factorise like import
```
var(
    a string
)
```

```
fmt.Printf("a= %v and its type %T ",a,a)
```
We can convert the above code
```
fmt.Printf("a= %v and its type %[1]T ",a)
```

### Types

- bool
- string
- int,int8,int16,int32,int64
- uint,uint8,uint16,uint32,uint64
- float32,float64

### Special types
- byte for int8
- rune for int32
- uintptr

### Type Conversion

```
T(v)
```

## IF condition
**Oneliner**
```
if intialization;condition{

}

``` 
### Loops
go has only one loop and its for
- no brackets around the three components

```
for n < 1000 {n+=n}
```
### Switch Statement

```
switch init;expression{
    case a:
    case b:
    default
}
```
>In go we don't have break instead we have the fallthrough which is the opposite

We use fall through if certain code to be repeated for many conditions

### Headless switch
```
switch{
    case i==2:
    case i==3
}
```
### Type declaration

```
type Distance int
```

To perform operation both has to be same type
- convert into the base type like int ,float
- convert into the any one type see eg

### Struct
```
type vector struct{
    x int
    y int
}
```

**Initilazation**
```
v1=:vector{1,2}
v2:=vector{x:1}
v3:=vector{}
```

### Pointers

& --> address
* --> dereference operator

> In struct we don't have to use the dereference operator
# Glossary
```
1. Hello -world
2. Import
4. function
5. variables
6. type conversion
7. if statement
8. for loop
10 . switch with fall through
11. type example
12. struct
13. Pointer to struct
```