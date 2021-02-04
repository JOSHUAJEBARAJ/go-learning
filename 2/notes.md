## Logics

## If statement

```
if <condition>{

}
```

## if else

```
if <condition>{

}else{ // both has be in the same line

}
```

## else if

```
if <condition>{

}else if <condition>{ // both has be in the same line

}
```
## The Initial if Statement or onliner if 
https://www.calhoun.io/one-liner-if-statements-with-errors/#:~:text=The%20first%20bit%20of%20the,a%20little%20easier%20to%20read.

```
if err := validate(input); err != nil {

    fmt.Println(err)

}
```
## Switch
```
switch <initial statement>; <expresion> {

case <expresion>:

  <statements>

case <expresion>, <expresion>:

  <statements>

default:

  <statements>

}
```
- both initial and expression are optional
- fallthrough command

## Mutliple case

```
switch <initial statement>; <expresion> {

case <expresion>:

  <statements>

case <expresion>, <expresion>:

  <statements>

default:

  <statements>

}
```

### Expressionless switch Statements

```
switch{ // value stored as the true
    case 

}
```
- case should return true or false
- The default switch value is true

### For Loop

```
for <initial statement>; <condition>; <post statement> {

  <statements>

}
```

Infinite loop

```
for{

}
```

### 2nd variant for loop

```
for <key>, <value> := range <map> {

  <statements>

}
```

> we can use _ to ignore key or value


Arrays and slices - for loop
maps - for range 
We can also use for range for array and slice