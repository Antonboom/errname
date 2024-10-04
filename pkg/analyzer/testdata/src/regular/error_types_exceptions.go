package regular

import "strings"

type ValidationErrors []string

func (ve ValidationErrors) Error() string { return strings.Join(ve, "\n") }

type validationErrors []string

func (ve validationErrors) Error() string { return strings.Join(ve, "\n") }

type TenErrors [10]string

func (te TenErrors) Error() string { return strings.Join(te[:], "\n") }

type tenErrors [10]string

func (te tenErrors) Error() string { return strings.Join(te[:], "\n") }

type MultiErr []error             // want "the error type name `MultiErr` should conform to the `XxxErrors` or `XxxError` format"
func (me MultiErr) Error() string { return "" }

type multiErr []error             // want "the error type name `multiErr` should conform to the `xxxErrors` or `xxxError` format"
func (me multiErr) Error() string { return "" }

type Twoerr [2]error            // want "the error type name `Twoerr` should conform to the `XxxErrors` or `XxxError` format"
func (te Twoerr) Error() string { return te[0].Error() + "\n" + te[1].Error() }

type twoErrorss [2]error            // want "the error type name `twoErrorss` should conform to the `xxxErrors` or `xxxError` format"
func (te twoErrorss) Error() string { return te[0].Error() + "\n" + te[1].Error() }

type MultiError []error

func (me MultiError) Error() string { return "" }

type multiError []error

func (me multiError) Error() string { return "" }

type TwoError [2]error

func (te TwoError) Error() string { return te[0].Error() + "\n" + te[1].Error() }

type twoError [2]error

func (te twoError) Error() string { return te[0].Error() + "\n" + te[1].Error() }
