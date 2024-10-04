package regular

import (
	"errors"
	"fmt"
)

var (
	EOF          = errors.New("end of file")
	ErrEndOfFile = errors.New("end of file")
	errEndOfFile = errors.New("end of file")

	EndOfFileError = errors.New("end of file") // want "the sentinel error name `EndOfFileError` should conform to the `ErrXxx` format"
	ErrorEndOfFile = errors.New("end of file") // want "the sentinel error name `ErrorEndOfFile` should conform to the `ErrXxx` format"
	EndOfFileErr   = errors.New("end of file") // want "the sentinel error name `EndOfFileErr` should conform to the `ErrXxx` format"
	endOfFileError = errors.New("end of file") // want "the sentinel error name `endOfFileError` should conform to the `errXxx` format"
	errorEndOfFile = errors.New("end of file") // want "the sentinel error name `errorEndOfFile` should conform to the `errXxx` format"
)

const maxSize = 256

var (
	ErrOutOfSize = fmt.Errorf("out of size (max %d)", maxSize)
	errOutOfSize = fmt.Errorf("out of size (max %d)", maxSize)

	OutOfSizeError = fmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error name `OutOfSizeError` should conform to the `ErrXxx` format"
	outOfSizeError = fmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error name `outOfSizeError` should conform to the `errXxx` format"
)

func errInsideFuncIsNotSentinel() error {
	var lastErr error
	return lastErr
}

var _ = func() {
	run("", func() {
		var orderErr error
		_ = orderErr
	})
}

func run(_ string, _ func()) {}
