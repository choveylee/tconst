package tconst

// Application error code constants describe application-layer outcomes. They
// complement, rather than replace, HTTP status codes and should be combined
// according to the response contract defined by each service.
const (
	// ErrorCodeSuccess indicates that no application error occurred.
	ErrorCodeSuccess = 0

	// ErrorCodeErrCodeIllegal indicates that the provided application error code
	// value is not recognized.
	ErrorCodeErrCodeIllegal = 1001
	// ErrorCodeRequestParamIllegal indicates that one or more request parameters
	// are invalid or disallowed.
	ErrorCodeRequestParamIllegal = 1002

	// ErrorCodeAccessTokenIllegal indicates that the access token is missing,
	// expired, or malformed.
	ErrorCodeAccessTokenIllegal = 1003
	// ErrorCodeRequestTimeIllegal indicates that the request timestamp is
	// missing, expired, or otherwise invalid under the service's time-skew
	// validation policy.
	ErrorCodeRequestTimeIllegal = 1004

	// ErrorCodeMysqlServerAbnormal indicates a failure in the MySQL backend.
	ErrorCodeMysqlServerAbnormal = 1101
	// ErrorCodeRedisServerAbnormal indicates a failure in the Redis backend.
	ErrorCodeRedisServerAbnormal = 1102
	// ErrorCodeEsServerAbnormal indicates a failure in the Elasticsearch
	// backend.
	ErrorCodeEsServerAbnormal = 1103

	// ErrorCodeUnknownServerAbnormal indicates an unclassified infrastructure
	// failure.
	ErrorCodeUnknownServerAbnormal = 1104

	// ErrorCodeActionIllegal indicates that the requested route or action is not
	// permitted.
	ErrorCodeActionIllegal = 1201
	// ErrorCodePermitCountIllegal indicates that a caller-specific or resource-
	// specific quota has been exceeded.
	ErrorCodePermitCountIllegal = 1202

	// ErrorCodeAuthRequestLimit indicates that the caller exceeded an
	// authentication or access rate limit.
	ErrorCodeAuthRequestLimit = 1301
)
