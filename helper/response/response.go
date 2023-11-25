package response

type ResponseWeb struct {
	Code    int         `json:"code"`
	Massage string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponese(code int, message string, data any) ResponseWeb {
	return ResponseWeb{
		Code:    code,
		Massage: message,
		Data:    data,
	}
}
