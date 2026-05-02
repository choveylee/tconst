# tconst

Shared numeric constants for HTTP response status codes and application error
codes intended for use across Go services.

## Requirements

- Go 1.25 or later

## Installation

```bash
go get github.com/choveylee/tconst
```

## Usage

```go
import "github.com/choveylee/tconst"

// HTTP response status code (for example, in a JSON API)
w.WriteHeader(tconst.StatusCodeOk)

// Application error code (for example, in a response payload)
code := tconst.ErrorCodeRequestParamIllegal
```

When a timestamp validation failure must be distinguished from a general bad
request, return `tconst.StatusCodeInvalidTime` together with
`tconst.ErrorCodeRequestTimeIllegal`. When a request is throttled, pair
`tconst.StatusCodeTooManyRequests` with `tconst.ErrorCodeAuthRequestLimit`.

## Constants

| Group | File | Description |
|-------|------|-------------|
| HTTP status | `status_code.go` | HTTP response status constants such as `StatusCodeOk`, `StatusCodeUnauthorized`, `StatusCodeTooManyRequests`, and `StatusCodeServerError`. |
| Application errors | `error_code.go` | Application error constants such as `ErrorCodeSuccess`, `ErrorCodeRequestTimeIllegal`, and the MySQL, Redis, and Elasticsearch infrastructure error codes. |

Some application conditions intentionally share the same HTTP status value. For
scenarios that require more precise client handling, return the appropriate
application error code alongside the HTTP status code.
