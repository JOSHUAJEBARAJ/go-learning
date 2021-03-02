## Http 

### URl 

https://www.google.com/

- protocol 
- hostname
- uri - additional path
- query parameters - separated by ?  and joined by & 
  
### Get request (3 steps)
- create the client
- read the response

```
ioutil.ReadAll(r.Body)
```
- convert into string


### Post  (3 steps) 
- create the body
- convert into byte
- send the request 

Post has 3 paramter
- data
- host
- type
  
## Custom client
- create the client 

```
	client := http.Client{Timeout: 11 * time.Second}
```
- create the new request
- create the custom header
- create the req
```
14.1 Get request
eg - get request
14.2 JSON request
eg-2 post
14.3 Post request
14.4 custom header
```