package helpers

import "net/http"

//ResponseHelper godoc
type ResponseHelper struct {
}

//ResponseHelperHandler godoc
func ResponseHelperHandler() ResponseHelper {
	return ResponseHelper{}
}

//Response godoc
type Response struct {
	Code    int         `json:"code"`
	Error   []string    `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

//SuccessResponse godoc
func (handler *ResponseHelper) SuccessResponse(data interface{}, message string) Response {
	response := Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: message,
		Status:  StatusSucces,
	}

	return response
}

//StatusBadRequest godoc
func (handler *ResponseHelper) StatusBadRequest(data interface{}, message string) Response {
	response := Response{
		Code:    http.StatusBadRequest,
		Data:    data,
		Message: message,
		Status:  StatusFailed,
	}

	return response
}

//InternalServerError godoc
func (handler *ResponseHelper) InternalServerError(data interface{}, message string) Response {
	response := Response{
		Code:    http.StatusInternalServerError,
		Data:    data,
		Message: message,
		Status:  StatusFailed,
	}

	return response
}
