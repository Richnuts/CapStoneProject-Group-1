package common

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessDataResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func CustomResponse(code int, status string, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"status":  status,
		"message": message,
	}
}

func ResponseDataSuccess(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":   200,
		"status": "Success",
		"data":   data,
	}
}

//NewInternalServerErrorResponse default internal server error response
func SuccessOperation(message string) DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
		message,
	}
}

//NewInternalServerErrorResponse default internal server error response
func InternalServerError() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
		"Servernya meleduk boss",
	}
}

//NewNotFoundResponse default not found error response
func NotFound() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
		"Data yang dicari tidak ada",
	}
}

//NewBadRequestResponse default not found error response
func BadRequest() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
		"ga jalan cuk",
	}
}

//ForbiddedRequest default not found error response
func ForbiddedRequest() DefaultResponse {
	return DefaultResponse{
		403,
		"Forbidded Request",
		"login dulu boss",
	}
}

//NewConflictResponse default not found error response
func Conflict() DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
		"data telah diubah",
	}
}

//NewUnauthorizedResponse default not found error response
func Unauthorized() DefaultResponse {
	return DefaultResponse{
		401,
		"Unauthorized Request",
		"yang bener pake tokennya mas",
	}
}
