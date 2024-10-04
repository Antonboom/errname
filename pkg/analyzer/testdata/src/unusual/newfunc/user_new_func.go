package newfunc

var FromUserNewError = new() // want "the sentinel error name `FromUserNewError` should conform to the `ErrXxx` format"

func new() error {
	return nil
}
