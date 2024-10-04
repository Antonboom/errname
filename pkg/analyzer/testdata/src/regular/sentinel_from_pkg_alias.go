package regular

import (
	stderrors "errors"
	stdfmt "fmt"
)

var ErrAlias = stderrors.New("err from alias")
var AliasErr = stderrors.New("err from alias") // want "the sentinel error name `AliasErr` should conform to the `ErrXxx` format"

var ErrOutOfSize2 = stdfmt.Errorf("out of size (max %d)", maxSize)
var OutOfSizeError2 = stdfmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error name `OutOfSizeError2` should conform to the `ErrXxx` format"
