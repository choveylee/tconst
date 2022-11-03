/**
 * @Author: lidonglin
 * @Description:
 * @File:  status_code.go
 * @Version: 1.0.0
 * @Date: 2022/11/03 10:43
 */

package tconst

const (
	StatusCodeTypeOk = 200 // 成功

	StatusCodeTypeInvalidReq = 400 // 请求缺少某个必需参数，包含一个不支持的参数或参数值，或者格式不正确
	StatusCodeTypeNotFound   = 404 // 请求失败，请求所希望得到的资源未被在服务器上发现。在参数相同的情况下，不应该重复请求
	StatusCodeTypeForbidden  = 403 // 用户没有对当前动作的权限，引导重新身份验证并不能提供任何帮助，而且这个请求也不应该被重复提交

	StatusCodeTypeServerError = 500 // 服务器出现异常情况 可稍等后重新尝试请求，但需有尝试上限，建议最多3次，如一直失败，则中断并告知用户

	StatusCodeTypeInvalidTime      = 400 // 客户端时间不正确，应请求服务器时间重新构造
	StatusCodeTypeAccessDenied     = 401 // AccessToken访问拒绝
	StatusCodeTypeMethodNotAllowed = 405 // 请求行中指定的请求方法不能被用于请求相应的资源
)
