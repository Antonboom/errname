package regular

import (
	"errors"
	"fmt"
)

var (
	EOF          = errors.New("end of file")
	ErrEndOfFile = errors.New("end of file")
	errEndOfFile = errors.New("end of file")

	EndOfFileError = errors.New("end of file") // want "the sentinel error `EndOfFileError` should be of the form ErrXxx"
	ErrorEndOfFile = errors.New("end of file") // want "the sentinel error `ErrorEndOfFile` should be of the form ErrXxx"
	EndOfFileErr   = errors.New("end of file") // want "the sentinel error `EndOfFileErr` should be of the form ErrXxx"
	endOfFileError = errors.New("end of file") // want "the sentinel error `endOfFileError` should be of the form errXxx"
	errorEndOfFile = errors.New("end of file") // want "the sentinel error `errorEndOfFile` should be of the form errXxx"
)

const maxSize = 256

var (
	ErrOutOfSize = fmt.Errorf("out of size (max %d)", maxSize)
	errOutOfSize = fmt.Errorf("out of size (max %d)", maxSize)

	OutOfSizeError = fmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error `OutOfSizeError` should be of the form ErrXxx"
	outOfSizeError = fmt.Errorf("out of size (max %d)", maxSize) // want "the sentinel error `outOfSizeError` should be of the form errXxx"
)

func errInsideFuncIsNotSentinel() error {
	var lastErr error
	return lastErr
}
