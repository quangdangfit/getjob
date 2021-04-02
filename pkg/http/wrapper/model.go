package wrapper

// Response body
type Response struct {
	HttpCode int         `json:"http_code"`
	Error    error       `json:"error"`
	Data     interface{} `json:"data"`
}
