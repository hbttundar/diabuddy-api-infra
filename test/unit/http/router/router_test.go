package router_test

import (
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockEngine struct {
	mock.Mock
}

type MockRouter struct {
	engine     *MockEngine
	middleware []router.Middleware
}

func NewMockRouter() *MockRouter {
	return &MockRouter{
		engine:     new(MockEngine),
		middleware: []router.Middleware{},
	}
}

func (m *MockRouter) GET(path string, handler router.RouteHandler) {
	m.engine.Called(path, mock.Anything) // ✅ don't compare handler directly
}

func (m *MockRouter) POST(path string, handler router.RouteHandler) {
	m.engine.Called(path, mock.Anything)
}

func (m *MockRouter) PUT(path string, handler router.RouteHandler) {
	m.engine.Called(path, mock.Anything)
}

func (m *MockRouter) PATCH(path string, handler router.RouteHandler) {
	m.engine.Called(path, mock.Anything)
}

func (m *MockRouter) DELETE(path string, handler router.RouteHandler) {
	m.engine.Called(path, mock.Anything)
}

func (m *MockRouter) Run(addr string) error {
	args := m.engine.Called(addr)
	return args.Error(0)
}

func TestMockRouter_DelegatesToEngine(t *testing.T) {
	mr := NewMockRouter()
	mr.engine.On("GET", "/ping", mock.Anything).Return()
	mr.engine.On("Run", ":8080").Return(nil)

	mr.GET("/ping", func() {})
	err := mr.Run(":8080")

	assert.NoError(t, err)
	mr.engine.AssertCalled(t, "GET", "/ping", mock.Anything)
	mr.engine.AssertCalled(t, "Run", ":8080")
}
