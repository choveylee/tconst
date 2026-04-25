# tconst

Shared constants for HTTP status codes and application error codes, for use across Go services.

## Requirements

- Go 1.25 or later

## Install

```bash
go get github.com/choveylee/tconst
```

## Usage

```go
import "github.com/choveylee/tconst"

// HTTP status (e.g. for JSON APIs)
w.WriteHeader(tconst.StatusCodeOk)

// Business error code (e.g. in response body)
code := tconst.ErrorCodeRequestParamIllegal
```

For timestamp validation failures, return `tconst.StatusCodeInvalidTime` with
`tconst.ErrorCodeRequestTimeIllegal`. For rate limiting, pair
`tconst.StatusCodeTooManyRequests` with `tconst.ErrorCodeAuthRequestLimit`.

## Constants

| Group | File | Description |
|-------|------|-------------|
| HTTP status | `status_code.go` | Numeric codes such as `StatusCodeOk`, `StatusCodeUnauthorized`, `StatusCodeTooManyRequests`, `StatusCodeServerError`. |
| Application errors | `error_code.go` | Numeric business error codes (e.g. `ErrorCodeSuccess`, `ErrorCodeRequestTimeIllegal`, MySQL/Redis/Elasticsearch infrastructure errors). |

Some HTTP-level distinctions share the same numeric status (for example, multiple `400` cases); use the business error code for fine-grained handling where needed.
