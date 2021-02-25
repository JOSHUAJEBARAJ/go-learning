## File System

### Flags and Arguments

The flag method takes three parameter
- name - name of the flag
- value - default value 
- Usage - the usage of flag
- return value - the value is pointer   

Use the program with - h to see the usage flag
flag.PrintDefaults() - to show the default value

## Signals
If we use system signals then defer will not  work

## Creating the files
- create function will take name and path as the argument
  
## Writing to the file
Write function - takes bytes of array as input
Write Sting function - takes string as input
We can use io.Writefile for creating and writing 

>It will truncate the file if exist
### Checking the file for existence

os.Stat() - will return the file info if it doesn't exist then it will multiple return error we have to check the check error type 
## Reading the Whole File at Once
There are two methods

- io.ReadFile - will read the whole file and returns the byte - the output has to be converted into the string
- ReadAll takes io.Reader - we have to take two steps

## Append file

we can use openfile method

```
Openfile(filename,flag,permission)
```
## File flag
1. os.O_CREATE - will create the file if not exist

```
os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
f.write([]newnotes)
```

### Reading CSV
strings.NewReader(in)

```
 for {

  record, err := r.Read()

  if err == io.EOF {

    break

  }

  if err != nil {

    log.Fatal(err)

  }

  fmt.Println(record)

  }
```
```
eg-1 flag example
eg-2 mandatory flag
eg-4 - creating the empty files
eg-5  - writing to the file
eg-6 - Creating and writing using ioutil writefile
eg-7 - checking the file existence
eg-8 Readfile
eg-9 ReadAll
12.2 - Backing up
eg-10 - reading csv
```