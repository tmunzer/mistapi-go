// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package framework_integrations

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// AppHandler is an interface that represents any HTTP handler compatible with the standard library.
type AppHandler interface {
	http.Handler
}

// FiberAdapter adapts a fiber.App to the standard http.Handler interface.
type FiberAdapter struct {
	app *fiber.App
}

// ServeHTTP implements the http.Handler interface for FiberAdapter.
// It allows a fiber.App to handle HTTP requests as a standard http.Handler.
func (f FiberAdapter) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request) {
	response, err := f.app.Test(request, -1)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	defer response.Body.Close()
	for key, values := range response.Header {
		for _, value := range values {
			writer.Header().Add(key, value)
		}
	}
	writer.WriteHeader(response.StatusCode)
	io.Copy(writer, response.Body)
}

// NewFiberApp creates a new Fiber application and registers the given event handler for POST /events.
// Returns an AppHandler compatible with the standard http.Handler interface.
func NewFiberApp(eventHandler func(*fiber.Ctx) error) AppHandler {
	app := fiber.New()
	app.Post("/events", eventHandler)
	return FiberAdapter{app: app}
}

// NewChiApp creates a new Chi application and registers the given event handler for POST /events.
// Returns an AppHandler compatible with the standard http.Handler interface.
func NewChiApp(eventHandler func(http.ResponseWriter, *http.Request)) AppHandler {
	app := chi.NewRouter()
	app.Post("/events", eventHandler)
	return app
}

// NewEchoApp creates a new Echo application and registers the given event handler for POST /events.
// Returns an AppHandler compatible with the standard http.Handler interface.
func NewEchoApp(eventHandler func(echo.Context) error) AppHandler {
	app := echo.New()
	app.POST("/events", eventHandler)
	return app.Server.Handler
}

// NewGinApp creates a new Gin application and registers the given event handler for POST /events.
// Returns an AppHandler compatible with the standard http.Handler interface.
func NewGinApp(eventHandler func(*gin.Context)) AppHandler {
	app := gin.New()
	app.POST("/events", eventHandler)
	return app
}
