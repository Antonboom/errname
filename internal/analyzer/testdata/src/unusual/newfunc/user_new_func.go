package newfunc

var FromUserNewError = new() // want "the sentinel error `FromUserNewError` should be of the form ErrXxx"

func new() error {
	return nil
}
