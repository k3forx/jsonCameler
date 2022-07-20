# jsonCameler

`jsonCameler` can be used to check whether JSON tags of structs are written in lowerCamel case.

## Installation

```bash
go install github.com/k3forx/jsonCameler/cmd/jsonCameler@latest
```

## How to use

```bash
go vet -vettool=`which jsonCameler` ./...
```

Here is an example.

```golang
package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`        // OK
	FirstName string    `json:"firstName"` // OK
	LastName  string    `json:"last-name"` // invalid
	BirthDay  time.Time `json:"-"`         // OK
}

func main() {
	...
}
```

The result will be like...

```bash
go vet -vettool=`which jsonCameler` ./
./main.go:16:2: invalid JSON tag last-name
```