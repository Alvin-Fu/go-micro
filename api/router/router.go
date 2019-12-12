// Package router provides api service routing
package router

import (
	"net/http"

	"github.com/micro/go-micro/api"
)

// Router is used to determine an endpoint for a request
// 路由器用于确定请求的终点
type Router interface {
	// Returns options
	Options() Options
	// Stop the router
	Close() error
	// Endpoint returns an api.Service endpoint or an error if it does not exist
	Endpoint(r *http.Request) (*api.Service, error)
	// Route returns an api.Service route
	Route(r *http.Request) (*api.Service, error)
}
