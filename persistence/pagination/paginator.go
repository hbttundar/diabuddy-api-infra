package pagination

// PaginationMeta represents metadata about paginated responses.
type PaginationMeta struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"` // Total items across all pages
	Last    int  `json:"last"`  // Last page number
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

// Paginator defines methods any pagination implementation should provide.
type Paginator interface {
	Page() int
	Limit() int
	Total() int
	Data() any
	First() int
	Last() int
	TotalPages() int
	HasNext() bool
	HasPrev() bool
	ToResponse() PaginatedResponse
}
