package domain

type Pagination struct {
	Items interface{} `json:"items"`
	Total int32       `json:"total"`
}
