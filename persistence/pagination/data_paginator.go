package pagination

import "math"

// PaginatedResponse is the final response structure returned to clients.
type PaginatedResponse struct {
	Data any `json:"data"`
	PaginationMeta
}

// DataPaginator implements the Paginator interface using core Pagination.
type DataPaginator struct {
	core *Pagination
}

func NewDataPaginator(p *Pagination) *DataPaginator {
	return &DataPaginator{core: p}
}

func (p *DataPaginator) Page() int  { return p.core.Page }
func (p *DataPaginator) Limit() int { return p.core.Limit }
func (p *DataPaginator) Total() int { return p.core.Total }
func (p *DataPaginator) Data() any  { return p.core.Data }
func (p *DataPaginator) First() int { return 1 }
func (p *DataPaginator) Last() int  { return p.TotalPages() }
func (p *DataPaginator) TotalPages() int {
	if p.core.Limit == 0 || p.core.Total == 0 {
		return 0
	}
	return int(math.Ceil(float64(p.core.Total) / float64(p.core.Limit)))
}
func (p *DataPaginator) HasNext() bool { return p.Page() < p.Last() }
func (p *DataPaginator) HasPrev() bool { return p.Page() > 1 }

func (p *DataPaginator) ToResponse() any {
	return PaginatedResponse{
		Data: p.Data(),
		PaginationMeta: PaginationMeta{
			Page:    p.Page(),
			Limit:   p.Limit(),
			Total:   p.Total(),
			Last:    p.Last(),
			HasNext: p.HasNext(),
			HasPrev: p.HasPrev(),
		},
	}
}
