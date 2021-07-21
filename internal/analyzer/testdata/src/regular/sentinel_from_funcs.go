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

	AErr = newSomeTypeWithPtr()     // want `the sentinel error "AErr" should be of the form ErrXxx`
	BErr = newSomeTypeWithPtr2()    // want `the sentinel error "BErr" should be of the form ErrXxx`
	CErr = newSomeTypeWithoutPtr()  // want `the sentinel error "CErr" should be of the form ErrXxx`
	DErr = newSomeTypeWithoutPtr2() // want `the sentinel error "DErr" should be of the form ErrXxx`
	EErr = new(SomeTypeWithPtr)     // want `the sentinel error "EErr" should be of the form ErrXxx`
	FErr = &SomeTypeWithPtr{}       // want `the sentinel error "FErr" should be of the form ErrXxx`
	GErr = SomeTypeWithoutPtr{}     // want `the sentinel error "GErr" should be of the form ErrXxx`

	AErrr error = newSomeTypeWithPtr2()   // want `the sentinel error "AErrr" should be of the form ErrXxx`
	BErrr error = newSomeTypeWithoutPtr() // want `the sentinel error "BErrr" should be of the form ErrXxx`
	CErrr error = new(SomeTypeWithPtr)    // want `the sentinel error "CErrr" should be of the form ErrXxx`
	DErrr error = &SomeTypeWithPtr{}      // want `the sentinel error "DErrr" should be of the form ErrXxx`
	EErrr error = SomeTypeWithoutPtr{}    // want `the sentinel error "EErrr" should be of the form ErrXxx`

	ErrByAnonymousFunc = func() error { return nil }
	ByAnonymousFuncErr = func() error { return io.EOF }() // want `the sentinel error "ByAnonymousFuncErr" should be of the form ErrXxx`
)

var (
	InitializedLaterError            error              // want `the sentinel error "InitializedLaterError" should be of the form ErrXxx`
	InitializedLaterImplicitError    SomeTypeWithoutPtr // want `the sentinel error "InitializedLaterImplicitError" should be of the form ErrXxx`
	InitializedLaterImplicitPtrError *SomeTypeWithPtr   // want `the sentinel error "InitializedLaterImplicitPtrError" should be of the form ErrXxx`
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
