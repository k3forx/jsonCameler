package main

import (
	"github.com/k3forx/jsonCameler"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(jsonCameler.Analyzer) }
