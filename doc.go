// Package tconst provides shared numeric constants for HTTP response status codes
// and application-level error codes used by Go services and internal APIs.
//
// # HTTP status codes
//
// StatusCode values follow conventional HTTP semantics as defined by RFC 9110.
// Distinct application conditions may intentionally share the same numeric
// status. When clients must distinguish between such conditions, services should
// include an application error code in the response. For example,
// StatusCodeInvalidTime and StatusCodeInvalidReq both map to 400, while
// ErrorCodeRequestTimeIllegal may be used to indicate a retryable clock-skew
// validation failure.
//
// # Error codes
//
// ErrorCode values identify application-level validation failures,
// infrastructure failures, routing violations, and rate-limiting conditions.
// They are independent of HTTP status codes unless a service contract defines
// an explicit mapping. Common mappings include ErrorCodeAccessTokenIllegal with
// StatusCodeUnauthorized and ErrorCodeAuthRequestLimit with
// StatusCodeTooManyRequests.
package tconst
