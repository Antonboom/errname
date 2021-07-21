package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Antonboom/errname/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.ErrName)
}
