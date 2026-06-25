package middleware

type APIResponse struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}

type ResponseInfo struct {
	Code       int         `json:"code"`
	Message    string      `json:"message,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
