## Packages
- Naming should be short and precise 
- Capital means exported non Capital means Un exported

## GOROOT and GOPATH

GOROOT tells the location of the standard packages

GOPATH tells the 3rd party library

GOPATH has 3 folders
- bin  - binary that are created
- src - actual source code
- pkg - stores the object files

```
$GOPATH/src/person/address/  - can be accessed by

import "person/address"
```

## Package Alias

```
import f "fmt"
```

## Main Package

THere are 2 types of packages
- executable
- non executable

Go build 
- create the binary inside the main packages and its name is same as folder name

## init function

These are used to set the initial value

Order of execution
```
imported package initialization
package variables
package initialization
main method
```