/**
 * @Author: lidonglin
 * @Description:
 * @File:  error_code.go
 * @Version: 1.0.0
 * @Date: 2022/11/03 10:43
 */

package tconst

const (
	ErrorCodeSuccess = 0 // 成功

	ErrorCodeErrCodeIllegal      = 1001 // 错误码非法
	ErrorCodeRequestParamIllegal = 1002 // 请求参数非法

	ErrorCodeAccessTokenIllegal = 1003 // AccessToken非法

	ErrorCodeDbServerAbnormal      = 1101 // Database服务器异常
	ErrorCodeRedServerAbnormal     = 1102 // Redis服务器异常
	ErrorCodeEsServerAbnormal      = 1103 // ElasticSearch服务器异常
	ErrorCodeUnknownServerAbnormal = 1104 // 未知服务器异常

	ErrorCodeActionIllegal      = 1201 // 访问路径非法
	ErrorCodePermitCountIllegal = 1202 // 访问数量上限

	ErrorCodeAuthRequestLimit = 1301 // 超过访问上限
)
