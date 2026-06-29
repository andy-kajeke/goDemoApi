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

type ResponsePagination struct {
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"totalPages"`
	HasNextPage bool  `json:"hasNextPage"`
}
