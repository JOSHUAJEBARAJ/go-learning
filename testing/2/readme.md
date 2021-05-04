## Naming Conventions

- [ ] Add _test to the filenames
- [ ] Prefix tests with **Test**
- [ ] Accept the Parameter - *testing.T
- [ ] We can use `package package_name` - to whitebox testing
- [ ] We can use `package package_name_test` - to perform black box


## Writing test

Immediate Failure - program exit
```
t.FaiilNow()
t.Fatal()
t.Fatalf()
``` 
Non-immediate failure

```
t.fail()
t.Error()
t.Errorf()
```

## Running test

```
go test

go test .

go test -v

 go test -run Greet // only the single test
```

## Coverage

```
go test -cover

go test -coverprofile cover.out

go tool cover -func cover.out

go tool cover -html cover.out 

```