package response

import "github.com/cloudwego/hertz/pkg/app"

type SuccessRes struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	ShowType     int         `json:"showType,omitempty"` // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	TraceId      string      `json:"traceId,omitempty"`  // Convenient for back-end Troubleshooting: unique request ID
	Host         string      `json:"host,omitempty"`     // onvenient for backend Troubleshooting: host of current access server

}

func Success(c *app.RequestContext, data interface{}) {
	c.JSON(200, &SuccessRes{
		Success: true,
		Data:    data,
	})
}
