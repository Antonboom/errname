package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrName(t *testing.T) {
	pkgs := []string{
		"regular",
		"unusual/errortype",
		"unusual/generics",
		"unusual/newfunc",
	}
	analysistest.Run(t, analysistest.TestData(), New(), pkgs...)
}
