## array

```
[<size>]<type>
```
We can initialize at the declaration
```
[5]string{9,9,9,9,9}
```
Another quick trick is we can ignore the length if we are intialize with the length
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
"jebaraj"
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