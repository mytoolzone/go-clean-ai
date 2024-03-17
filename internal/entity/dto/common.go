package dto

type Page struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type Pagination struct {
	TotalCount int64 `json:"total_count"`
	HasMore    bool  `json:"has_more"`
	Page
}
