// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// RedisTestResult defines model for RedisTestResult.
type RedisTestResult struct {
	RMQConsumers   *int      `json:"RMQConsumers,omitempty"`
	RMQFaultQueues *[]string `json:"RMQFaultQueues,omitempty"`
	RMQLongQueues  *[]string `json:"RMQLongQueues,omitempty"`
	RMQQueues      *int      `json:"RMQQueues,omitempty"`
	Name           string    `json:"name"`
	Result         bool      `json:"result"`
}

// TestResult defines model for TestResult.
type TestResult struct {
	Name   string `json:"name"`
	Result bool   `json:"result"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Description of the endpoint
	// (GET /api/v1/{env}/{identity}/run-test/rmq)
	GetApiV1EnvIdentityRunTestRmq(ctx echo.Context, env string, identity string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiV1EnvIdentityRunTestRmq converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiV1EnvIdentityRunTestRmq(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "env" -------------
	var env string

	err = runtime.BindStyledParameterWithLocation("simple", false, "env", runtime.ParamLocationPath, ctx.Param("env"), &env)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter env: %s", err))
	}

	// ------------- Path parameter "identity" -------------
	var identity string

	err = runtime.BindStyledParameterWithLocation("simple", false, "identity", runtime.ParamLocationPath, ctx.Param("identity"), &identity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter identity: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiV1EnvIdentityRunTestRmq(ctx, env, identity)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/:env/:identity/run-test/rmq", wrapper.GetApiV1EnvIdentityRunTestRmq)

}