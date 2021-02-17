## array

```
[<size>]<type>
```
We can initialize at the declaration
```
[5]string{9,9,9,9,9}
```
Another quick trick is we can ignore the length if we are intialise with the length
```
[...]string{9,9,9,9,9}
```
> But we cannot change the length

## Comparing array

Same length and size will can be compared

> slice and map, are not comparable 

## Initializing Arrays Using Keys

Storing the value at particular index
```
 arr2 := [...]int{9: 0} // storing the value at the particular index
```
## Reading from an Array

```
func message() string{

    arr := [...]string{

"joshua",
"jebaraj",
    }
    return fmt.Sprintln(arr[0],arr[1])
}
```
> Comma for the end of the array too

## Looping an Array

While passing use the same format
```
arr [10] int
```
> If we loop over in the for range it will create the loop

## Slice

- Wrapper around the array
- Same as one type
- **append** function to increase the length of slice
- We should slice as much as possible

### Declaration
```
var args []string 
```
### append
```
var [] string args
args = append(args, value)
```

> remember argument and type for the function definition

### Appending Multiple Items to a Slice


### Creating Slices from Slices and Arrays

```
[<low>:<high>]
```

> It doesnt include the last value 

- if we omit the low value it will start from 0
- if we omit the high value it will end at the last element
  
  ## Understanding Slice Internals

  Slice have 3 properties
- length
- pointer to the array
- starting point of the array

## Using make
make(<sliceType>, <length>, <capacity>).


## Maps

```
map[<key_type>]<value_type>
```

> , Go returns the zero value for the map's value type. 


### Reading from Map

```
<value>, <exists_value> := <map>[<key>].
```

### Deleting an Element from a Map

```
delete(map,id)
```

## Type Conversions

- Smaller to larger no Issue
- Larger to smaller will cause issue like overflow
- Usigned int to signed int also cause overflow

```
<type>(<value>)
int8(value)
```

## Type assertion

- fmt take interface has the input
- interface has no fields and functions

```
<value>.(<type>)
<value>, <ok> := <value>.(type)
```

## Type Switch

```
switch <value> := <value>.(type) {

case <type>:

  <statement>

case <type>, <type>:

  <statement>

default:

  <statement>

}
```

## Matching float type
```
case float32, float64:

  if f, ok := t.(float64); ok {

    return fmt.Sprint(f * 2), nil

  }
  return fmt.Sprint(t.(float32) * 2), nil
```