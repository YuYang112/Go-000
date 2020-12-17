package services

type Page struct {
	Total   int64       `json:"total"`
	PerPage int         `json:"per_page"`
	PageNum int         `json:"page_num"`
	Data    interface{} `json:"data"`
}
