package regular

import "net"

var (
	InvalidAddrError       = new(net.AddrError) // want "the sentinel error name `InvalidAddrError` should conform to the `ErrXxx` format"
	InvalidAddrErr   error = new(net.AddrError) // want "the sentinel error name `InvalidAddrErr` should conform to the `ErrXxx` format"
	NotErr                 = new(NotErrorType)

	Aa  = new(someTypeWithPtr) // want "the sentinel error name `Aa` should conform to the `ErrXxx` format"
	Bb  = someTypeWithoutPtr{} // want "the sentinel error name `Bb` should conform to the `ErrXxx` format"
	Bbb = someTypeWithPtr{}

	cC error = new(someTypeWithPtr) // want "the sentinel error name `cC` should conform to the `errXxx` format"
	dD error = someTypeWithoutPtr{} // want "the sentinel error name `dD` should conform to the `errXxx` format"

	Alias = InvalidAddrErr // want "the sentinel error name `Alias` should conform to the `ErrXxx` format"
)
