package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		next = middleware(next)
	}

	for _, middleware := range mngr.globalMiddlewares {
		next = middleware(next)
	}
	return next
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {

	for _, middleware := range mngr.globalMiddlewares {
		next = middleware(next)
	}
	return next
}
