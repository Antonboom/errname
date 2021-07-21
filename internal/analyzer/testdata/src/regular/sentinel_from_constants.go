package regular

type constError string

func (e constError) Error() string {
	return string(e)
}

const (
	ErrTooManyErrors constError = "too many errors found"

	ErrorTooMany1 constError = "too many errors found"             // want `the sentinel error "ErrorTooMany1" should be of the form ErrXxx`
	ErrorTooMany2            = constError("too many errors found") // want `the sentinel error "ErrorTooMany2" should be of the form ErrXxx`
)
