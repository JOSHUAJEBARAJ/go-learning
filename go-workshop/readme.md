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

## Data Structure

### Array

- have fixed length
```
[n] T
```

```
var a[10]string
```
### Slice

- Slice pointer to an array
- has length and capacity
- zero value is **nil**

We can use **make** function

```
a:=make([]int,0)
```

we can use range

```
s[low:high]
```

> high =high-1

### Append

Append will create the new array and store it in the variable
> to append the slice(no element) to existing slice use ...


### for range

we can either ignore key or value

*To ignore key*
```
for _,key:= range slice{

}
```

*To ignore value*

>We dont have to use the _
```
for i:=range slice{

}
```

### Maps

Maps are like an key value pair

We can use make to create the map

```
m=make(map[int]string)
```

#### Assigning the value

```
m[key]=vale
```

#### Deleting the value

```
delete(variable,key)
```
#### Getting value
There are 2 types
- elem=m[key]
- elem,ok=m[key]

# OOPS

## Methods
Go doesn't have class
```
func (v vertex) name() return type{}

```

### Methods vs Function
```
v := Vertex{1, 3}
a := v.Abs()

v := Vertex{1, 3}
a := Abs(v)

```

> remember to use the * next to the type what annie said

We can use methods only in the same package

```
func (f float64) IsPositive() bool {
	return f >= 0
}

```
The above wont work since float belong to another package
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
14. Array
15. Slice
16. for range
17. Map
19. Methods
```