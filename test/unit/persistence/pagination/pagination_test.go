package pagination_test

import (
	"github.com/hbttundar/diabuddy-api-infra/persistence/pagination"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestNewPagination_WithOptions(t *testing.T) {
	p := pagination.NewPagination(
		pagination.WithPage(2),
		pagination.WithLimit(20),
		pagination.WithTotal(100),
		pagination.WithData([]string{"a", "b"}),
	)

	assert.Equal(t, 2, p.Page)
	assert.Equal(t, 20, p.Limit)
	assert.Equal(t, 100, p.Total)
	assert.Equal(t, []string{"a", "b"}, p.Data)
}

func TestWithRequest_ValidParams(t *testing.T) {
	r := &http.Request{URL: &url.URL{RawQuery: "page=3&limit=15"}}
	p := pagination.NewPagination(pagination.WithRequest(r))

	assert.Equal(t, 3, p.Page)
	assert.Equal(t, 15, p.Limit)
}

func TestWithRequest_InvalidParams(t *testing.T) {
	r := &http.Request{URL: &url.URL{RawQuery: "page=abc&limit=-5"}}
	p := pagination.NewPagination(pagination.WithRequest(r))

	assert.Equal(t, 1, p.Page)   // fallback default
	assert.Equal(t, 10, p.Limit) // fallback default
}

func TestDataPaginator_ToResponse(t *testing.T) {
	p := pagination.NewPagination(
		pagination.WithPage(2),
		pagination.WithLimit(10),
		pagination.WithTotal(45),
		pagination.WithData([]int{1, 2, 3}),
	)
	dp := pagination.NewDataPaginator(p)
	r := dp.ToResponse()

	assert.Equal(t, 2, r.Page)
	assert.Equal(t, 10, r.Limit)
	assert.Equal(t, 45, r.Total)
	assert.Equal(t, 5, r.Last)
	assert.True(t, r.HasNext)
	assert.True(t, r.HasPrev)
	assert.Equal(t, []int{1, 2, 3}, r.Data)
}
