package wrapper

type ErrorWrapper struct {
	Code   string  `json:"code"`
	Msg    string  `json:"msg"`
	Detail *string `json:"detail,omitempty"`
}
