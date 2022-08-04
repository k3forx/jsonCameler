# jsonCameler [![Go Report Card](https://goreportcard.com/badge/github.com/k3forx/jsonCameler)](https://goreportcard.com/report/github.com/k3forx/jsonCameler)

`jsonCameler` can be used to check whether JSON tags of structs are written in lowerCamel case. You can skip lint check by adding comment `nocaml` to the fields of the struct or to the above line of the declaration of the struct.

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
	ID        int       `json:"id"`        
	FirstName string    `json:"firstName"` 
	LastName  string    `json:"last-name"` // invalid
	Username  string    `json:"user-name"` // nocamel
	BirthDay  time.Time `json:"-"`         
}

// nocamel
type Score struct {
	ID        int
	UserID    int       `json:"user_id"`
	Score     int       `json:"s-co-re"`
	CreatedAt time.Time `json:"created_at" some tags`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func main() {
	...
}
```

The result will be like...

```bash
go vet -vettool=`which jsonCameler` ./
./main.go:16:2: invalid JSON tag `last-name`
```