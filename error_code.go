package tconst

// Application error codes identify business-layer outcomes. They are not a substitute
// for HTTP status; combine them according to each API's contract.
const (
	// ErrorCodeSuccess indicates that no application error occurred.
	ErrorCodeSuccess = 0

	// ErrorCodeErrCodeIllegal indicates that the supplied error code value is not recognized.
	ErrorCodeErrCodeIllegal = 1001
	// ErrorCodeRequestParamIllegal indicates invalid or disallowed request parameters.
	ErrorCodeRequestParamIllegal = 1002

	// ErrorCodeAccessTokenIllegal indicates that the access token is missing, expired, or malformed.
	ErrorCodeAccessTokenIllegal = 1003
	// ErrorCodeRequestTimeIllegal indicates that the request timestamp is missing,
	// expired, or otherwise invalid for time-skew validation.
	ErrorCodeRequestTimeIllegal = 1004

	// ErrorCodeMysqlServerAbnormal indicates a MySQL backend failure.
	ErrorCodeMysqlServerAbnormal = 1101
	// ErrorCodeRedisServerAbnormal indicates a Redis backend failure.
	ErrorCodeRedisServerAbnormal = 1102
	// ErrorCodeEsServerAbnormal indicates an Elasticsearch backend failure.
	ErrorCodeEsServerAbnormal = 1103

	// ErrorCodeUnknownServerAbnormal indicates an infrastructure failure not classified above.
	ErrorCodeUnknownServerAbnormal = 1104

	// ErrorCodeActionIllegal indicates that the requested path or action is not allowed.
	ErrorCodeActionIllegal = 1201
	// ErrorCodePermitCountIllegal indicates that a per-caller or resource quota was exceeded.
	ErrorCodePermitCountIllegal = 1202

	// ErrorCodeAuthRequestLimit indicates that the caller exceeded an authentication or access rate limit.
	ErrorCodeAuthRequestLimit = 1301
)
