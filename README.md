# errname

![Latest release](https://img.shields.io/github/v/release/Antonboom/errname)
[![CI](https://github.com/Antonboom/errname/actions/workflows/ci.yml/badge.svg)](https://github.com/Antonboom/errname/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Antonboom/errname)](https://goreportcard.com/report/github.com/Antonboom/errname)
[![Coverage](https://coveralls.io/repos/github/Antonboom/errname/badge.svg?branch=master)](https://coveralls.io/github/Antonboom/errname?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Checks that sentinel errors are prefixed with the `Err` and error types
are suffixed with the `Error` or `Errors`.

## Installation & usage

```
$ go install github.com/Antonboom/errname@latest
$ errname ./...
```

## Motivation

[The convention](https://go.dev/wiki/Errors#naming) states that
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

### Why not [revive/error-naming](https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md#error-naming)?

At the time of writing this linter, I was unaware of **revive**. <br>
**errname** performs more complex and better checks anyway.

## Examples

```go
// Bad.
type DecodeErr struct{} // the error type name `DecodeErr` should conform to the `xxxError` format
func (d *DecodeErr) Error() string { /*...*/ }

// Good.
type DecodeError struct{}
func (d *DecodeError) Error() string { /*...*/ }
```

```go
// Bad.
var InvalidURLErr = errors.New("invalid url") // the sentinel error name `InvalidURLErr` should conform to the `ErrXxx` format 

// Good.
var ErrInvalidURL = errors.New("invalid url") // or errInvalidURL
```

More examples in [tests](https://github.com/Antonboom/errname/blob/master/pkg/analyzer/facts_test.go).

## Assumptions

<details>
  <summary>Click to expand</summary>

<br>

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

- Linter only checks the correctness of the suffix and prefix and their **uniqueness**. The logical meaning of the
  identifier remains on the developer's conscience:

```go
// Bad.
var ErrExecErr = errors.New("exec query error")

// Good.
var ErrExecQuery = errors.New("exec query error")
var ErrGdfjnskjdfskf = errors.New("strange error") // on the developer's conscience
```

- For error types over array/slice the `Errors` suffix is also allowed:

```go
// Bad.
type ValidationErrs []string
func (ve ValidationErrs) Error() string { /*...*/ }

// Good.
type ValidationErrors []string
func (ve ValidationErrors) Error() string { /*...*/ }
```

</details>
        
## Large projects examples

<details>
  <summary>Golang source code</summary>

```go
$ errname ./src/...
go/src/runtime/error.go:72:6: the error type name `errorString` should conform to the `xxxError` format
go/src/runtime/error.go:80:6: the error type name `errorAddressString` should conform to the `xxxError` format
go/src/runtime/panic.go:180:5: the sentinel error name `shiftError` should conform to the `errXxx` format
go/src/runtime/panic.go:187:5: the sentinel error name `divideError` should conform to the `errXxx` format
go/src/runtime/panic.go:194:5: the sentinel error name `overflowError` should conform to the `errXxx` format
go/src/runtime/panic.go:201:5: the sentinel error name `floatError` should conform to the `errXxx` format
go/src/runtime/panic.go:208:5: the sentinel error name `memoryError` should conform to the `errXxx` format
go/src/errors/errors.go:63:6: the error type name `errorString` should conform to the `xxxError` format
go/src/math/bits/bits_errors.go:12:5: the sentinel error name `overflowError` should conform to the `errXxx` format
go/src/math/bits/bits_errors.go:15:5: the sentinel error name `divideError` should conform to the `errXxx` format
go/src/syscall/syscall_unix.go:114:6: the error type name `Errno` should conform to the `XxxError` format
go/src/time/format.go:394:5: the sentinel error name `atoiError` should conform to the `errXxx` format
go/src/time/zoneinfo_read.go:110:5: the sentinel error name `badData` should conform to the `errXxx` format
go/src/io/fs/walk.go:15:5: the sentinel error name `SkipDir` should conform to the `ErrXxx` format
go/src/fmt/scan.go:465:5: the sentinel error name `complexError` should conform to the `errXxx` format
go/src/fmt/scan.go:466:5: the sentinel error name `boolError` should conform to the `errXxx` format
go/src/archive/tar/common.go:39:6: the error type name `headerError` should conform to the `xxxError` format
go/src/context/context.go:157:5: the sentinel error name `Canceled` should conform to the `ErrXxx` format
go/src/context/context.go:161:5: the sentinel error name `DeadlineExceeded` should conform to the `ErrXxx` format
go/src/math/big/float.go:77:6: the error type name `ErrNaN` should conform to the `XxxError` format
go/src/crypto/x509/internal/macos/security.go:39:6: the error type name `OSStatus` should conform to the `XxxError` format
go/src/net/cgo_unix.go:34:6: the error type name `addrinfoErrno` should conform to the `xxxError` format
go/src/crypto/x509/x509.go:875:6: the error type name `UnhandledCriticalExtension` should conform to the `XxxError` format
go/src/crypto/x509/pem_decrypt.go:110:5: the sentinel error name `IncorrectPasswordError` should conform to the `ErrXxx` format
go/src/crypto/x509/root.go:18:2: the sentinel error name `systemRootsErr` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:18:2: the sentinel error name `alertCloseNotify` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:19:2: the sentinel error name `alertUnexpectedMessage` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:20:2: the sentinel error name `alertBadRecordMAC` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:21:2: the sentinel error name `alertDecryptionFailed` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:22:2: the sentinel error name `alertRecordOverflow` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:23:2: the sentinel error name `alertDecompressionFailure` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:24:2: the sentinel error name `alertHandshakeFailure` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:25:2: the sentinel error name `alertBadCertificate` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:26:2: the sentinel error name `alertUnsupportedCertificate` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:27:2: the sentinel error name `alertCertificateRevoked` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:28:2: the sentinel error name `alertCertificateExpired` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:29:2: the sentinel error name `alertCertificateUnknown` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:30:2: the sentinel error name `alertIllegalParameter` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:31:2: the sentinel error name `alertUnknownCA` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:32:2: the sentinel error name `alertAccessDenied` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:33:2: the sentinel error name `alertDecodeError` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:34:2: the sentinel error name `alertDecryptError` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:35:2: the sentinel error name `alertExportRestriction` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:36:2: the sentinel error name `alertProtocolVersion` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:37:2: the sentinel error name `alertInsufficientSecurity` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:38:2: the sentinel error name `alertInternalError` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:39:2: the sentinel error name `alertInappropriateFallback` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:40:2: the sentinel error name `alertUserCanceled` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:41:2: the sentinel error name `alertNoRenegotiation` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:42:2: the sentinel error name `alertMissingExtension` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:43:2: the sentinel error name `alertUnsupportedExtension` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:44:2: the sentinel error name `alertCertificateUnobtainable` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:45:2: the sentinel error name `alertUnrecognizedName` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:46:2: the sentinel error name `alertBadCertificateStatusResponse` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:47:2: the sentinel error name `alertBadCertificateHashValue` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:48:2: the sentinel error name `alertUnknownPSKIdentity` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:49:2: the sentinel error name `alertCertificateRequired` should conform to the `errXxx` format
go/src/crypto/tls/alert.go:50:2: the sentinel error name `alertNoApplicationProtocol` should conform to the `errXxx` format
go/src/path/filepath/path.go:337:5: the sentinel error name `SkipDir` should conform to the `ErrXxx` format
go/src/net/http/h2_bundle.go:1016:5: the sentinel error name `http2errReadEmpty` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1212:2: the sentinel error name `http2errMixPseudoHeaderTypes` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1213:2: the sentinel error name `http2errPseudoAfterRegular` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1712:5: the sentinel error name `http2ErrFrameTooLarge` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1866:2: the sentinel error name `http2errStreamID` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1867:2: the sentinel error name `http2errDepStreamID` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1868:2: the sentinel error name `http2errPadLength` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:1869:2: the sentinel error name `http2errPadBytes` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3400:5: the sentinel error name `http2errTimeout` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3519:5: the sentinel error name `http2errClosedPipeWrite` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3629:2: the sentinel error name `http2errClientDisconnected` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3630:2: the sentinel error name `http2errClosedBody` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3631:2: the sentinel error name `http2errHandlerComplete` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:3632:2: the sentinel error name `http2errStreamClosed` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:4526:5: the sentinel error name `http2errPrefaceTimeout` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:4746:5: the sentinel error name `http2errHandlerPanicked` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:6287:2: the sentinel error name `http2ErrRecursivePush` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:6288:2: the sentinel error name `http2ErrPushLimitReached` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:6930:5: the sentinel error name `http2ErrNoCachedConn` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7016:2: the sentinel error name `http2errClientConnClosed` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7017:2: the sentinel error name `http2errClientConnUnusable` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7018:2: the sentinel error name `http2errClientConnGotGoAway` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7471:5: the sentinel error name `http2errRequestCanceled` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7803:2: the sentinel error name `http2errStopReqBodyWrite` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7806:2: the sentinel error name `http2errStopReqBodyWriteAndCancel` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:7808:2: the sentinel error name `http2errReqBodyTooLong` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:8667:5: the sentinel error name `http2errClosedResponseBody` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:9021:2: the sentinel error name `http2errResponseHeaderListSize` should conform to the `errXxx` format
go/src/net/http/h2_bundle.go:9022:2: the sentinel error name `http2errRequestHeaderListSize` should conform to the `errXxx` format
go/src/go/scanner/errors.go:37:6: the error type name `ErrorList` should conform to the `XxxError` format
go/src/html/template/template.go:34:5: the sentinel error name `escapeOK` should conform to the `errXxx` format
go/src/image/png/reader.go:128:5: the sentinel error name `chunkOrderError` should conform to the `errXxx` format
go/src/bufio/scan_test.go:308:5: the sentinel error name `testError` should conform to the `errXxx` format
go/src/crypto/tls/handshake_client_test.go:1993:5: the sentinel error name `brokenConnErr` should conform to the `errXxx` format
go/src/database/sql/sql_test.go:4281:5: the sentinel error name `pingError` should conform to the `errXxx` format
go/src/errors/wrap_test.go:216:6: the error type name `errorT` should conform to the `xxxError` format
go/src/errors/wrap_test.go:229:6: the error type name `errorUncomparable` should conform to the `xxxError` format
go/src/fmt/errors_test.go:75:6: the error type name `errString` should conform to the `xxxError` format
go/src/html/template/exec_test.go:233:5: the sentinel error name `myError` should conform to the `errXxx` format
go/src/html/template/exec_test.go:1313:5: the sentinel error name `alwaysError` should conform to the `errXxx` format
go/src/net/http/transport_test.go:6280:5: the sentinel error name `timeoutProtoErr` should conform to the `errXxx` format
go/src/text/template/exec_test.go:229:5: the sentinel error name `myError` should conform to the `errXxx` format
go/src/text/template/exec_test.go:1305:5: the sentinel error name `alwaysError` should conform to the `errXxx` format
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
terraform/internal/getmodules/file_detector.go:59:6: the error type name `MaybeRelativePathErr` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:13:6: the error type name `ErrHostNoProviders` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:39:6: the error type name `ErrHostUnreachable` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:57:6: the error type name `ErrUnauthorized` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:80:6: the error type name `ErrProviderNotFound` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:104:6: the error type name `ErrRegistryProviderNotKnown` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:123:6: the error type name `ErrPlatformNotSupported` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:159:6: the error type name `ErrProtocolNotSupported` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:181:6: the error type name `ErrQueryFailed` should conform to the `XxxError` format
terraform/internal/getproviders/errors.go:219:6: the error type name `ErrRequestCanceled` should conform to the `XxxError` format
terraform/internal/registry/errors.go:10:6: the error type name `errModuleNotFound` should conform to the `xxxError` format
terraform/internal/backend/remote-state/consul/client.go:36:5: the sentinel error name `lostLockErr` should conform to the `errXxx` format
terraform/internal/command/cliconfig/credentials.go:408:6: the error type name `ErrUnwritableHostCredentials` should conform to the `XxxError` format
```

</details>
