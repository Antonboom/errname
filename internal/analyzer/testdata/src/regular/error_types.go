package regular

type NotErrorType struct{}

func (t NotErrorType) Set() {}
func (t NotErrorType) Get() {}

type DNSConfigError struct{}

func (D DNSConfigError) Error() string { return "DNS config error" }

type someTypeWithoutPtr struct{}           // want `the error type "someTypeWithoutPtr" should be of the form xxxError`
func (s someTypeWithoutPtr) Error() string { return "someTypeWithoutPtr" }

type SomeTypeWithoutPtr struct{}           // want `the error type "SomeTypeWithoutPtr" should be of the form XxxError`
func (s SomeTypeWithoutPtr) Error() string { return "SomeTypeWithoutPtr" }

type someTypeWithPtr struct{}            // want `the error type "someTypeWithPtr" should be of the form xxxError`
func (s *someTypeWithPtr) Error() string { return "someTypeWithPtr" }

type SomeTypeWithPtr struct{}            // want `the error type "SomeTypeWithPtr" should be of the form XxxError`
func (s *SomeTypeWithPtr) Error() string { return "SomeTypeWithPtr" }
