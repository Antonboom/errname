package regular

import (
	stderrors "errors"
	stdfmt "fmt"
)

var (
	ErrAlias = stderrors.New("err from alias")
	AliasErr = stderrors.New("err from alias") // want "the sentinel error `AliasErr` should be of the form ErrXxx"
)

var (
	ErrOutOfSize2   = stdfmt.Errorf("out of size (max %d)", maxSize)
	OutOfSizeError2 = stdfmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error `OutOfSizeError2` should be of the form ErrXxx"
)
