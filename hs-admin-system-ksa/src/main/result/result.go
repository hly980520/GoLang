package result

type ApiResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Error 返回业务错误提醒
func Error(code string, message string) (apiResult ApiResult) {
	var result ApiResult
	result.Code = code
	result.Message = message
	return result
}

// Ok 返回请求成功接口
func Ok(data any) (apiResult ApiResult) {
	var result ApiResult
	result.Code = "10000000"
	result.Data = data
	return result
}
