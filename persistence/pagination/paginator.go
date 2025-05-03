package pagination

type PaginationMeta struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	Last    int  `json:"last"`
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

type Paginator interface {
	Page() int
	Limit() int
	Total() int
	Data() any

	TotalPages() int
	HasNext() bool
	HasPrev() bool
	Last() int
	First() int

	ToResponse() any
}
