// Package tconst defines shared numeric constants for HTTP response status codes
// and application-level error codes for use in APIs and internal services.
//
// # HTTP status codes
//
// StatusCode values follow conventional HTTP semantics (see RFC 9110). Several
// distinct application conditions may share the same numeric status; responses
// should include an application error code when finer-grained handling is required.
// For example, StatusCodeInvalidTime and StatusCodeInvalidReq are both 400, and
// callers should pair clock-skew cases with ErrorCodeRequestTimeIllegal when
// they need retryable behavior.
//
// # Error codes
//
// ErrorCode values identify validation failures, infrastructure faults, routing
// issues, and rate limits. They are orthogonal to HTTP status unless combined
// by a service's response contract. Common mappings include
// ErrorCodeAccessTokenIllegal with StatusCodeUnauthorized and
// ErrorCodeAuthRequestLimit with StatusCodeTooManyRequests.
package tconst
