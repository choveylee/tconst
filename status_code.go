package tconst

// HTTP status code constants for API responses, listed in ascending numeric
// order. Multiple exported names may intentionally share the same numeric value
// when distinct application conditions map to the same HTTP status.
const (
	// StatusCodeOk indicates that the request completed successfully.
	StatusCodeOk = 200

	// StatusCodeInvalidReq indicates that the request is invalid, such as when
	// required fields are missing, input is malformed, or parameter values are
	// unsupported.
	StatusCodeInvalidReq = 400
	// StatusCodeInvalidTime indicates that the client timestamp failed
	// validation and should be corrected before the request is retried.
	StatusCodeInvalidTime = 400

	// StatusCodeUnauthorized indicates that authentication failed or the access
	// token was rejected.
	StatusCodeUnauthorized = 401

	// StatusCodeForbidden indicates that the authenticated caller is not
	// authorized to perform the requested action.
	StatusCodeForbidden = 403

	// StatusCodeNotFound indicates that the requested resource could not be
	// found.
	StatusCodeNotFound = 404

	// StatusCodeMethodNotAllowed indicates that the HTTP method is not supported
	// for the requested resource.
	StatusCodeMethodNotAllowed = 405
	// StatusCodeTooManyRequests indicates that the caller exceeded a rate limit
	// and should retry only after an appropriate backoff interval.
	StatusCodeTooManyRequests = 429

	// StatusCodeServerError indicates an unexpected internal server failure.
	StatusCodeServerError = 500
)
