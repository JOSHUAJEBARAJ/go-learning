## JSON

- key value pair

There are 4 values
- string
- int 
- null
- array

Object = struct

## Decoding JSON

```
json --> go type
```
- we can use the **encoding/json**
- unmarshal packages two values  json data (slice) , struct pointer 

> The ` symbol is a backtick and not a single quote. It is used for string literals.
> To be able to unmarshal into a struct, the struct field must be exportable.The struct's field name must be capitalized.


```
d:=[]byte(`
	{
		"name":"joshua"
	}
`)

```



### Struct Tags
- We can use the struct tag to map the struct field to the json field
- The struct tag in the below format `key: "value"`
- valid:=json.Valid(d)  - to check whether it is valid json or not

Unmarshalling order
```
struct tag --> key match --> insensitive key   
```

### Encoding JSON
- golang to json
- we can use marshal indent to print pretty json.MarshalIndent(s,""," ")
### omit empty

omit empty to omit the value if the value of struct is empty

**Advanced omit empty**

```
`json:"id,omitempty"` - omit if empty or take the id as the Key name
`json:",omitempty"` - omit if empty or take the field name as the Key name
`json:",-"` - doesn't do marshalling (uses iphen not underscore)
```

### Unknown JSON Structures
- we can use map[string]interface{}

```
map[string]  = key
interface = value
```
- order is not gurranted
```
eg-1 unmarhsalling example
eg-2 struct tag
11.1 - un marshalling
eg-3 Marshalling
eg-4 Marhalling with tag
eg-5 omitempty
eg-6 advanced omit empty
eg-7 pretty print
11.2 - Marsalling
eg-8 unknown json
11.3 un marshalling with type
```