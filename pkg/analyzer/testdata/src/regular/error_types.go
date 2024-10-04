package regular

type NotErrorType struct{}

func (t NotErrorType) Set() {}
func (t NotErrorType) Get() {}

type DNSConfigError struct{}

func (D DNSConfigError) Error() string { return "DNS config error" }

type someTypeWithoutPtr struct{}           // want "the error type name `someTypeWithoutPtr` should conform to the `xxxError` format"
func (s someTypeWithoutPtr) Error() string { return "someTypeWithoutPtr" }

type SomeTypeWithoutPtr struct{}           // want "the error type name `SomeTypeWithoutPtr` should conform to the `XxxError` format"
func (s SomeTypeWithoutPtr) Error() string { return "SomeTypeWithoutPtr" }

type someTypeWithPtr struct{}            // want "the error type name `someTypeWithPtr` should conform to the `xxxError` format"
func (s *someTypeWithPtr) Error() string { return "someTypeWithPtr" }

type (
	SomeTypeAlias = SomeTypeWithPtr // want "the error type name `SomeTypeAlias` should conform to the `XxxError` format"

	SomeTypeWithPtr struct{} // want "the error type name `SomeTypeWithPtr` should conform to the `XxxError` format"
)

func (s *SomeTypeWithPtr) Error() string { return "SomeTypeWithPtr" }

type timeoutErr struct { // want "the error type name `timeoutErr` should conform to the `xxxError` format"
	error
}

type DeadlineErr struct { // want "the error type name `DeadlineErr` should conform to the `XxxError` format"
	timeoutErr
}
