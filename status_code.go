package tconst

// HTTP status code constants for API responses, ordered by ascending numeric value.
// Multiple symbols may share a value when distinct conditions map to the same HTTP status.
const (
	// StatusCodeOk indicates that the request completed successfully.
	StatusCodeOk = 200

	// StatusCodeInvalidReq indicates a bad request: missing required fields,
	// unsupported parameters or values, or malformed input.
	StatusCodeInvalidReq = 400
	// StatusCodeInvalidTime indicates that the client clock is invalid; the client
	// should obtain server time and retry with a corrected timestamp.
	StatusCodeInvalidTime = 400

	// StatusCodeAccessDenied indicates that authentication failed or the access token was rejected.
	StatusCodeAccessDenied = 401

	// StatusCodeForbidden indicates that the authenticated caller is not permitted
	// to perform this action; re-authentication will not resolve the denial.
	StatusCodeForbidden = 403

	// StatusCodeNotFound indicates that the target resource does not exist.
	// Identical requests must not be retried indefinitely when the resource remains absent.
	StatusCodeNotFound = 404

	// StatusCodeMethodNotAllowed indicates that the HTTP method is not allowed for the resource.
	StatusCodeMethodNotAllowed = 405

	// StatusCodeServerError indicates an unexpected server failure. Clients may retry
	// with bounded attempts (for example, up to three); persistent failure should surface to the user.
	StatusCodeServerError = 500
)
