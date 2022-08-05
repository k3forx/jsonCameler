package main

import (
	"github.com/k3forx/jsonCameler"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		jsonCameler.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
