package response

import "github.com/labstack/echo/v4"

type SuccessResp struct {
	successCode int
	Data        interface{}
}

type ErrorResp struct {
	errorCode int
	messages  interface{}
}

func SuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	resp := &SuccessResp{
		successCode: statusCode,
		Data:        data,
	}
	c.Response().WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(statusCode, resp, "  ")
}

func ErrorResponse(c echo.Context, errorCode int, messages interface{}) error {
	resp := &ErrorResp{
		errorCode: errorCode,
		messages:  messages,
	}
	c.Response().WriteHeader(errorCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(errorCode, resp, "  ")
}
