// Package tconst defines shared numeric constants for HTTP response status codes
// and application-level error codes for use in APIs and internal services.
//
// # HTTP status codes
//
// StatusCode values follow conventional HTTP semantics (see RFC 9110). Several
// distinct application conditions may share the same numeric status; responses
// should include an application error code when finer-grained handling is required.
//
// # Error codes
//
// ErrorCode values identify validation failures, infrastructure faults, routing
// issues, and rate limits. They are orthogonal to HTTP status unless combined
// by a service's response contract.
package tconst
