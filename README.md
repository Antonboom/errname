# errname

[![CI](https://github.com/Antonboom/errname/actions/workflows/ci.yml/badge.svg)](https://github.com/Antonboom/errname/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Antonboom/errname)](https://goreportcard.com/report/github.com/Antonboom/errname)
[![Coverage](https://coveralls.io/repos/github/Antonboom/errname/badge.svg?branch=master)](https://coveralls.io/github/Antonboom/errname?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

The [Golang](https://golang.org/) linter that checks that sentinel errors are prefixed with the `Err` and error types
are suffixed with the `Error`.

## Installation & usage

```
$ go get github.com/Antonboom/errname
$ errname ./...
```

## Motivation

[The convention](https://github.com/golang/go/wiki/Errors#naming) states that
> Error types end in "Error" and error variables start with "Err" or "err".

This can be found in the standard Go library:

```go
type AddrError struct{ /* ... */ }
type DecodeError struct{ /*...*/ }
type NoMatchingVersionError struct{ /*...*/ }
type nothingWrittenError struct{ /*...*/ }
type pasteIndicatorError struct{ /*...*/ }
type wrapError struct{ /*...*/ }

var ErrFinalToken = errors.New("final token")
var ErrRange = errors.New("value out of range")
var ErrUnsupportedAlgorithm = errors.New("x509: cannot verify signature: algorithm unimplemented")
var errMissing = errors.New("cannot find package")
var errNUL = errors.New("unexpected NUL in input")
var errTagValueSpace = errors.New("suspicious space in struct tag value")
```

Also, this can be found in some articles about errors in Go, for
example, [here](https://travix.io/errors-derived-from-constants-in-go-fda6748b4072):
> At first sight you’d think the naming is a bit weird, but that is because of conventions: types are suffixed with
> “Error”, while constants are prefixed with “Err”. So we’re just following the conventions here.

This is a good rule to improve the consistency of your code. **More uniformity, what could be better?**

### What if I think it's bullshit?

Just don't enable the linter.

## Examples

```go
// Bad.
type DecodeErr struct{}            // the error type "DecodeErr" should be of the form xxxError
func (d *DecodeErr) Error() string { /*...*/ }

// Good.
type DecodeError struct{}
func (d *DecodeError) Error() string { /*...*/ }
```

```go
// Bad.
var InvalidURLErr = errors.New("invalid url") // the sentinel error "InvalidURLErr" should be of the form ErrXxx

// Good.
var ErrInvalidURL = errors.New("invalid url") // or errInvalidURL
```

More examples in [tests](https://github.com/Antonboom/errname/blob/master/internal/analyzer/facts_test.go).

## Assumptions

- Only package level variables (and constants) are checked.
- Initialisms are ignored. As a result, all identifiers in a single case are ignored:

```go
var EOF = errors.New("end of file")
var eof = errors.New("end of file")
var W = errors.New("single letter error")   // on the developer's conscience
var ovflw = errors.New("value is overflow") // on the developer's conscience
```

- The naming of error constructors is not checked. But I recommend keeping it consistent:

```go
type DecodeError struct{}
func (d *DecodeError) Error() string { /*...*/ }

// Bad.
func NewErrDecode() error {
	return &DecodeError{}
}

// Good.
func NewDecodeError() error {
	return &DecodeError{}
}
```

- Package aliases are not supported if the source package and its directory differ in name.

- Not supported sentinel errors that were not created by a `errors`/`fmt` package and that do not have an explicit
  type `error`:

```go
var ErrUnsupported = new(net.AddrError)
var ErrSupported error = new(net.AddrError)
```

- Linter only checks the correctness of the suffix and prefix and their **uniqueness**. The logical meaning of the
  identifier remains on the developer's conscience.

```go
// Bad.
var ErrExecErr = errors.New("exec query error")

// Good.
var ErrExecQuery = errors.New("exec query error")
var ErrGdfjnskjdfskf = errors.New("strange error") // on the developer's conscience
```

- For error types over array/slice the `Errors` suffix is expected:

```go
// Bad.
type ValidationErrs []string
func (ve ValidationErrs) Error() string { /*...*/ }

// Good.
type ValidationErrors []string
func (ve ValidationErrors) Error() string { /*...*/ }
```

## Large projects examples

<details>
  <summary>Golang source code</summary>

```go
$ errname./src/...
/usr/local/go /src/runtime/error.go:72:6: the error type "errorString" should be of the form xxxError
/usr/local/go /src/runtime/error.go:80:6: the error type "errorAddressString" should be of the form xxxError
/usr/local/go /src/runtime/panic.go:180:5: the sentinel error "shiftError" should be of the form errXxx
/usr/local/go /src/runtime/panic.go:187:5: the sentinel error "divideError" should be of the form errXxx
/usr/local/go /src/runtime/panic.go:194:5: the sentinel error "overflowError" should be of the form errXxx
/usr/local/go /src/runtime/panic.go:201:5: the sentinel error "floatError" should be of the form errXxx
/usr/local/go /src/runtime/panic.go:208:5: the sentinel error "memoryError" should be of the form errXxx
/usr/local/go /src/errors/errors.go:63:6: the error type "errorString" should be of the form xxxError
/usr/local/go /src/math/bits/bits_errors.go:12:5: the sentinel error "overflowError" should be of the form errXxx
/usr/local/go /src/math/bits/bits_errors.go:15:5: the sentinel error "divideError" should be of the form errXxx
/usr/local/go /src/syscall/syscall_unix.go:114:6: the error type "Errno" should be of the form XxxError
/usr/local/go /src/time/format.go:394:5: the sentinel error "atoiError" should be of the form errXxx
/usr/local/go /src/time/zoneinfo_read.go:110:5: the sentinel error "badData" should be of the form errXxx
/usr/local/go /src/io/fs/walk.go:15:5: the sentinel error "SkipDir" should be of the form ErrXxx
/usr/local/go /src/fmt/scan.go:465:5: the sentinel error "complexError" should be of the form errXxx
/usr/local/go /src/fmt/scan.go:466:5: the sentinel error "boolError" should be of the form errXxx
/usr/local/go /src/archive/tar/common.go:39:6: the error type "headerError" should be of the form xxxError
/usr/local/go/src/context/context.go:157:5: the sentinel error "Canceled" should be of the form ErrXxx
/usr/local/go/src/context/context.go:161:5: the sentinel error "DeadlineExceeded" should be of the form ErrXxx
/usr/local/go/src/math/big/float.go:77:6: the error type "ErrNaN" should be of the form XxxError
/usr/local/go /src/crypto/x509/internal/macos/security.go:39:6: the error type "OSStatus" should be of the form XxxError
/usr/local/go /src/net/cgo_unix.go:34:6: the error type "addrinfoErrno" should be of the form xxxError
/usr/local/go /src/crypto/x509/x509.go:875:6: the error type "UnhandledCriticalExtension" should be of the form XxxError
/usr/local/go/src/crypto/x509/pem_decrypt.go:110:5: the sentinel error "IncorrectPasswordError" should be of the form ErrXxx
/usr/local/go /src/crypto/x509/root.go:18:2: the sentinel error "systemRootsErr" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:18:2: the sentinel error "alertCloseNotify" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:19:2: the sentinel error "alertUnexpectedMessage" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:20:2: the sentinel error "alertBadRecordMAC" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:21:2: the sentinel error "alertDecryptionFailed" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:22:2: the sentinel error "alertRecordOverflow" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:23:2: the sentinel error "alertDecompressionFailure" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:24:2: the sentinel error "alertHandshakeFailure" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:25:2: the sentinel error "alertBadCertificate" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:26:2: the sentinel error "alertUnsupportedCertificate" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:27:2: the sentinel error "alertCertificateRevoked" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:28:2: the sentinel error "alertCertificateExpired" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:29:2: the sentinel error "alertCertificateUnknown" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:30:2: the sentinel error "alertIllegalParameter" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:31:2: the sentinel error "alertUnknownCA" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:32:2: the sentinel error "alertAccessDenied" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:33:2: the sentinel error "alertDecodeError" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:34:2: the sentinel error "alertDecryptError" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:35:2: the sentinel error "alertExportRestriction" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:36:2: the sentinel error "alertProtocolVersion" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:37:2: the sentinel error "alertInsufficientSecurity" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:38:2: the sentinel error "alertInternalError" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:39:2: the sentinel error "alertInappropriateFallback" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:40:2: the sentinel error "alertUserCanceled" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:41:2: the sentinel error "alertNoRenegotiation" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:42:2: the sentinel error "alertMissingExtension" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:43:2: the sentinel error "alertUnsupportedExtension" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:44:2: the sentinel error "alertCertificateUnobtainable" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:45:2: the sentinel error "alertUnrecognizedName" should be of the form errXxx
/usr/local/go/src/crypto/tls/alert.go:46:2: the sentinel error "alertBadCertificateStatusResponse" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:47:2: the sentinel error "alertBadCertificateHashValue" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:48:2: the sentinel error "alertUnknownPSKIdentity" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:49:2: the sentinel error "alertCertificateRequired" should be of the form errXxx
/usr/local/go /src/crypto/tls/alert.go:50:2: the sentinel error "alertNoApplicationProtocol" should be of the form errXxx
/usr/local/go/src/path/filepath/path.go:337:5: the sentinel error "SkipDir" should be of the form ErrXxx
/usr/local/go /src/net/http/h2_bundle.go:1016:5: the sentinel error "http2errReadEmpty" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1212:2: the sentinel error "http2errMixPseudoHeaderTypes" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1213:2: the sentinel error "http2errPseudoAfterRegular" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1712:5: the sentinel error "http2ErrFrameTooLarge" should be of the form errXxx
/usr/local/go/src/net/http/h2_bundle.go:1866:2: the sentinel error "http2errStreamID" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1867:2: the sentinel error "http2errDepStreamID" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1868:2: the sentinel error "http2errPadLength" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:1869:2: the sentinel error "http2errPadBytes" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:3400:5: the sentinel error "http2errTimeout" should be of the form errXxx
/usr/local/go/src/net/http/h2_bundle.go:3519:5: the sentinel error "http2errClosedPipeWrite" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:3629:2: the sentinel error "http2errClientDisconnected" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:3630:2: the sentinel error "http2errClosedBody" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:3631:2: the sentinel error "http2errHandlerComplete" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:3632:2: the sentinel error "http2errStreamClosed" should be of the form errXxx
/usr/local/go/src/net/http/h2_bundle.go:4526:5: the sentinel error "http2errPrefaceTimeout" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:4746:5: the sentinel error "http2errHandlerPanicked" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:6287:2: the sentinel error "http2ErrRecursivePush" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:6288:2: the sentinel error "http2ErrPushLimitReached" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:6930:5: the sentinel error "http2ErrNoCachedConn" should be of the form errXxx
/usr/local/go/src/net/http/h2_bundle.go:7016:2: the sentinel error "http2errClientConnClosed" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:7017:2: the sentinel error "http2errClientConnUnusable" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:7018:2: the sentinel error "http2errClientConnGotGoAway" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:7471:5: the sentinel error "http2errRequestCanceled" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:7803:2: the sentinel error "http2errStopReqBodyWrite" should be of the form errXxx
/usr/local/go/src/net/http/h2_bundle.go:7806:2: the sentinel error "http2errStopReqBodyWriteAndCancel" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:7808:2: the sentinel error "http2errReqBodyTooLong" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:8667:5: the sentinel error "http2errClosedResponseBody" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:9021:2: the sentinel error "http2errResponseHeaderListSize" should be of the form errXxx
/usr/local/go /src/net/http/h2_bundle.go:9022:2: the sentinel error "http2errRequestHeaderListSize" should be of the form errXxx
/usr/local/go/src/go /scanner/errors.go:37:6: the error type "ErrorList" should be of the form XxxError
/usr/local/go /src/html/template/template.go:34:5: the sentinel error "escapeOK" should be of the form errXxx
/usr/local/go /src/image/png/reader.go:128:5: the sentinel error "chunkOrderError" should be of the form errXxx
/usr/local/go /src/bufio/scan_test.go:308:5: the sentinel error "testError" should be of the form errXxx
/usr/local/go /src/crypto/tls/handshake_client_test.go:1993:5: the sentinel error "brokenConnErr" should be of the form errXxx
/usr/local/go /src/database/sql/sql_test.go:4281:5: the sentinel error "pingError" should be of the form errXxx
/usr/local/go/src/errors/wrap_test.go:216:6: the error type "errorT" should be of the form xxxError
/usr/local/go/src/errors/wrap_test.go:229:6: the error type "errorUncomparable" should be of the form xxxError
/usr/local/go/src/fmt/errors_test.go:75:6: the error type "errString" should be of the form xxxError
/usr/local/go/src/html/template/exec_test.go:233:5: the sentinel error "myError" should be of the form errXxx
/usr/local/go /src/html/template/exec_test.go:1313:5: the sentinel error "alwaysError" should be of the form errXxx
/usr/local/go /src/net/http/transport_test.go:6280:5: the sentinel error "timeoutProtoErr" should be of the form errXxx
/usr/local/go /src/text/template/exec_test.go:229:5: the sentinel error "myError" should be of the form errXxx
/usr/local/go /src/text/template/exec_test.go:1305:5: the sentinel error "alwaysError" should be of the form errXxx

```

</details>

<details>
  <summary>Traefik</summary>

```go
$ errname./...
# no issues
```

</details>

<details>
  <summary>Terraform</summary>

```go
$ errname./...
terraform/internal/getmodules/file_detector.go:59:6: the error type "MaybeRelativePathErr" should be of the form XxxError
terraform/internal/getproviders/errors.go:13:6: the error type "ErrHostNoProviders" should be of the form XxxError
terraform/internal/getproviders/errors.go:39:6: the error type "ErrHostUnreachable" should be of the form XxxError
terraform/internal/getproviders/errors.go:57:6: the error type "ErrUnauthorized" should be of the form XxxError
terraform/internal/getproviders/errors.go:80:6: the error type "ErrProviderNotFound" should be of the form XxxError
terraform/internal/getproviders/errors.go:104:6: the error type "ErrRegistryProviderNotKnown" should be of the form XxxError
terraform/internal/getproviders/errors.go:123:6: the error type "ErrPlatformNotSupported" should be of the form XxxError
terraform/internal/getproviders/errors.go:159:6: the error type "ErrProtocolNotSupported" should be of the form XxxError
terraform/internal/getproviders/errors.go:181:6: the error type "ErrQueryFailed" should be of the form XxxError
terraform/internal/getproviders/errors.go:219:6: the error type "ErrRequestCanceled" should be of the form XxxError
terraform/internal/registry/errors.go:10:6: the error type "errModuleNotFound" should be of the form xxxError
terraform/internal/backend/remote-state/consul/client.go:36:5: the sentinel error "lostLockErr" should be of the form errXxx
terraform/internal/command/cliconfig/credentials.go:408:6: the error type "ErrUnwritableHostCredentials" should be of the form XxxError
```

</details>
