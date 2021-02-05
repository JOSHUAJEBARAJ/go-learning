## Notes

### True or false
0  means false 1 means true


## Numbers
- whole number 
- Floating number

Integer can be classified based on the below thing
- whether they can store negative number or not
- How big they can store

Types that can store negative numbers are called **signed integers**.

### Special types
- uint
- int

Both are either 32 bit or 64 bit
>An int on a 64-bit system is not an int64.

### Floating point
- float32
- float64

## Overflow and Wraparound

overflow this will happen when we allocate the number more than the space
```
var a unit=199
```

Wraparound means the number goes from its highest possible value to its smallest possible value
```
127,0,1,2
```

## big number
used to create the big number than the original type

## Byte
- alias of uint8 of 8 bit
- bit is either true or false
- 256 combination
  
## Text   

String can be defined in two ways

```
" " - this way
`` - this way
```

**raw** method is used when we specifying the directory

```
comment1 := `In "Windows" the user directory is "C:\Users\"`
```

## Rune

- Alias of int32
- 32 bit
- ASCII 128 code  can fit into 8 bit
- type with enough  single UTF-8 multi-byte character

> we can also use range to deal with rune

The below method is safely working with string
```
for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}}
```

###  nil

empty maps ,pointers,interfaces