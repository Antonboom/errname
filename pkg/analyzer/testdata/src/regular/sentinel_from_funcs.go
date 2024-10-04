package regular

import "io"

var (
	ErrA = newSomeTypeWithPtr()
	ErrB = newSomeTypeWithPtr2()
	ErrC = newSomeTypeWithoutPtr()
	ErrD = newSomeTypeWithoutPtr2()
	ErrE = new(SomeTypeWithPtr)
	ErrF = &SomeTypeWithPtr{}
	ErrG = SomeTypeWithoutPtr{}

	AErr = newSomeTypeWithPtr()     // want "the sentinel error name `AErr` should conform to the `ErrXxx` format"
	BErr = newSomeTypeWithPtr2()    // want "the sentinel error name `BErr` should conform to the `ErrXxx` format"
	CErr = newSomeTypeWithoutPtr()  // want "the sentinel error name `CErr` should conform to the `ErrXxx` format"
	DErr = newSomeTypeWithoutPtr2() // want "the sentinel error name `DErr` should conform to the `ErrXxx` format"
	EErr = new(SomeTypeWithPtr)     // want "the sentinel error name `EErr` should conform to the `ErrXxx` format"
	FErr = &SomeTypeWithPtr{}       // want "the sentinel error name `FErr` should conform to the `ErrXxx` format"
	GErr = SomeTypeWithoutPtr{}     // want "the sentinel error name `GErr` should conform to the `ErrXxx` format"

	AErrr error = newSomeTypeWithPtr2()   // want "the sentinel error name `AErrr` should conform to the `ErrXxx` format"
	BErrr error = newSomeTypeWithoutPtr() // want "the sentinel error name `BErrr` should conform to the `ErrXxx` format"
	CErrr error = new(SomeTypeWithPtr)    // want "the sentinel error name `CErrr` should conform to the `ErrXxx` format"
	DErrr error = &SomeTypeWithPtr{}      // want "the sentinel error name `DErrr` should conform to the `ErrXxx` format"
	EErrr error = SomeTypeWithoutPtr{}    // want "the sentinel error name `EErrr` should conform to the `ErrXxx` format"

	ErrByAnonymousFunc = func() error { return nil }
	ByAnonymousFuncErr = func() error { return io.EOF }() // want "the sentinel error name `ByAnonymousFuncErr` should conform to the `ErrXxx` format"
)

var (
	InitializedLaterError            error              // want "the sentinel error name `InitializedLaterError` should conform to the `ErrXxx` format"
	InitializedLaterImplicitError    SomeTypeWithoutPtr // want "the sentinel error name `InitializedLaterImplicitError` should conform to the `ErrXxx` format"
	InitializedLaterImplicitPtrError *SomeTypeWithPtr   // want "the sentinel error name `InitializedLaterImplicitPtrError` should conform to the `ErrXxx` format"
)

func init() {
	InitializedLaterError = newSomeTypeWithPtr()
	InitializedLaterImplicitError = newSomeTypeWithoutPtr()
	InitializedLaterImplicitPtrError = newSomeTypeWithPtr2()
}

func newSomeTypeWithPtr() error {
	return new(SomeTypeWithPtr)
}

func newSomeTypeWithPtr2() *SomeTypeWithPtr {
	return nil
}

func newSomeTypeWithoutPtr() SomeTypeWithoutPtr {
	return SomeTypeWithoutPtr{}
}

func newSomeTypeWithoutPtr2() error {
	return SomeTypeWithoutPtr{}
}
