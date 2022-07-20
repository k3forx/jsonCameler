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