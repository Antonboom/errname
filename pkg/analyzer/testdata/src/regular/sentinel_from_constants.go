package regular

type constError string

func (e constError) Error() string {
	return string(e)
}

const (
	ErrTooManyErrors constError = "too many errors found"

	ErrorTooMany1 constError = "too many errors found"             // want "the sentinel error name `ErrorTooMany1` should conform to the `ErrXxx` format"
	ErrorTooMany2            = constError("too many errors found") // want "the sentinel error name `ErrorTooMany2` should conform to the `ErrXxx` format"
)
