## Variables and operators

### Variables 

Variables can be declared in the many ways

- using var keyword

> Note doesn't use : 

```
var name string ="Joshua"
```

- Declaring Multiple Variables at Once with var

> Use bracket same as the import
```
var ()
```

- Skipping the Type or Value When Declaring Variables

```
var (
    age int 
    name= joshua
)
```

### Type Inference Gone Wrong

```
package main

import "math/rand"

func main() {

  var seed = 1234456789 // treated as int but expected int164

  rand.Seed(seed)

}
```

Solution to the above problem

```
var seed int64 = 1234456789
```

- using short hand

```
name:="Joshua"
```

- Declaring Multiple Variables from a Function (Returning multiple variable in the function)

```
func fullname () (string,string){
  return "joshua","jebaraj"
}
```

> Non-English Variable Names also possible

## operators

### Short hand

```
--: Reduce a number by 1
++: Increase a number by 1
+=: Add and assign
-=: Subtract and assign
```

## Comparing value

```
== True if two values are the same
!= True if two values are not the same
< True if the left value is less than the right value
<= True if the left value is less or equal to the right value
> True if the left value is greater than the right value
>= True if the left value is greater than or equal to the right value
```

## Logical operator
```
&& True if the left and right values are both true
|| True if one or both the left and right values are true
! This operator only works with a single value and results in true if the value is false
```

## Zero values

```
bool - false
number - 0
func,interfaces,slices,maps,pointers - nil
string - ""
```

## Substitution
```
%v - value
%T - type
%s - string
%d - number
%+v - only values 
%#v - additional value with name 
```

> + show value and variable
## Value vs pointer
Passing by value vs passing by reference
Pass by value not efficient , Go use **stack** Pointer uses **heap**

Pointer can be created in the four ways

Only star after type
```
var <name> *<type>
```
- using var
- using new keyword
- using existing variable
- using the  other data -type like struct

> While getting the value from the pointer prefix the variable with *

and pass panna 

## Getting a Value from a Pointer

Using the  dereference operator
```
fmt.Println(*<val>)
```
## Constant

using const keyword

```
constant <name> <type> = <value>

constant (

  <name1> <type1> = <value1>

  <name2> <type2> = <value3>

  <nameN> <typeN> = <valueN>

)
```

## Extras

### Enums
Using itoa
https://yourbasic.org/golang/iota/
### Scopes