package regular

import "net"

var (
	InvalidAddrError       = new(net.AddrError) // unsupported
	InvalidAddrErr   error = new(net.AddrError) // want "the sentinel error `InvalidAddrErr` should be of the form ErrXxx"
	NotErr                 = new(NotErrorType)

	Aa = new(someTypeWithPtr) // want "the sentinel error `Aa` should be of the form ErrXxx"
	Bb = someTypeWithoutPtr{} // want "the sentinel error `Bb` should be of the form ErrXxx"

	cC error = new(someTypeWithPtr) // want "the sentinel error `cC` should be of the form errXxx"
	dD error = someTypeWithoutPtr{} // want "the sentinel error `dD` should be of the form errXxx"
)
