package common

type Object interface {
}

type HttpResult struct {
	successful int
	errorCode  string
	errorMsg   string
	statusCode int
	data       Object
}

func BuildSuccHttpResult(data Object) *HttpResult {
	var result = new(HttpResult)
	result.successful = 1
	result.data = data
	result.statusCode = 200
	result.errorMsg = "success"
	return result
}

func BuildFailHttpResult(statusCode int, errorCode string, errorMsg string) *HttpResult {
	var result = new(HttpResult)
	result.successful = 0
	result.data = map[string]string{"errorCode": errorCode, "errorMsg": errorMsg}
	result.statusCode = statusCode
	result.errorMsg = errorMsg
	result.errorCode = errorCode
	return result
}
