package collection

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	TotalItem  int         `json:"total_items,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
