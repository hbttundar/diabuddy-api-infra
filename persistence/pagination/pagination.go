package pagination

import (
	"net/http"
	"strconv"
)

// Pagination holds paginated result metadata and data itself.
type Pagination struct {
	Data  any
	Page  int
	Limit int
	Total int
}

// Option defines a functional option for Pagination.
type Option func(*Pagination)

// NewPagination creates a new pagination instance with default or overridden settings.
func NewPagination(opts ...Option) *Pagination {
	p := &Pagination{
		Page:  1,
		Limit: 10,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithPage(page int) Option {
	return func(p *Pagination) {
		if page > 0 {
			p.Page = page
		}
	}
}

func WithLimit(limit int) Option {
	return func(p *Pagination) {
		if limit > 0 {
			p.Limit = limit
		}
	}
}

func WithTotal(total int) Option {
	return func(p *Pagination) {
		if total >= 0 {
			p.Total = total
		}
	}
}

func WithData(data any) Option {
	return func(p *Pagination) {
		p.Data = data
	}
}

// WithRequest parses page and limit from HTTP request query parameters.
func WithRequest(r *http.Request) Option {
	return func(p *Pagination) {
		q := r.URL.Query()
		if v, err := strconv.Atoi(q.Get("page")); err == nil && v > 0 {
			p.Page = v
		}
		if v, err := strconv.Atoi(q.Get("limit")); err == nil && v > 0 {
			p.Limit = v
		}
	}
}
