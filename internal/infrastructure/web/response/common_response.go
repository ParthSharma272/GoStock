package response

// StandardErrorResponse defines a standard error response structure.
type StandardErrorResponse struct {
	Error string `json:"error"`
}

// StandardSuccessResponse defines a standard success message response.
type StandardSuccessResponse struct {
	Message string `json:"message"`
}

// PaginatedResponse helps structure paginated data.
type PaginatedResponse struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	LastPage int64       `json:"last_page"`
}
