// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package webhooks_integrations

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"mistapi"
	"mistapi/events/webhooks"
	"net/http"
)

// WebhooksGinEventHandler receives and processes incoming events
func WebhooksGinEventHandler(c *gin.Context) {
	handler := webhooks.NewWebhooksHandler()
	parsingResult := handler.ParseEvent(c.Request)
	if event, ok := parsingResult.AsAlarms(); ok {
		c.JSON(200, map[string]any{
			"status":    "success",
			"eventInfo": fmt.Sprintf("Alarms event received %v", event),
		})
	} else {
		c.JSON(400, map[string]any{
			"status": "failure",
		})
	}
}

// WebhooksFiberEventHandler receives and processes incoming events
func WebhooksFiberEventHandler(c *fiber.Ctx) error {
	fastReq := c.Request()
	req, err := mistapi.FromFastHttpRequest(
		fastReq.Header.Method(),
		fastReq.URI().String(),
		fastReq.Body(),
		fastReq.Header.VisitAll,
		fastReq.URI().RequestURI(),
		fastReq.Header.VisitAllCookie,
	)
	if err != nil {
		return c.Status(500).JSON(map[string]any{
			"internal server error": err.Error(),
		})
	}

	handler := webhooks.NewWebhooksHandler()
	parsingResult := handler.ParseEvent(req)
	if event, ok := parsingResult.AsAlarms(); ok {
		return c.Status(200).JSON(map[string]any{
			"status":    "success",
			"eventInfo": fmt.Sprintf("Alarms event received %v", event),
		})
	} else {
		return c.Status(400).JSON(map[string]any{
			"status": "failure",
		})
	}
}

// WebhooksChiEventHandler receives and processes incoming events
func WebhooksChiEventHandler(
	writer http.ResponseWriter,
	request *http.Request) {
	handler := webhooks.NewWebhooksHandler()
	parsingResult := handler.ParseEvent(request)
	if event, ok := parsingResult.AsAlarms(); ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(map[string]any{
			"status":    "success",
			"eventInfo": fmt.Sprintf("Alarms event received %v", event),
		})
	} else {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(map[string]any{
			"status": "failure",
		})
	}
}

// WebhooksEchoEventHandler receives and processes incoming events
func WebhooksEchoEventHandler(c echo.Context) error {
	handler := webhooks.NewWebhooksHandler()
	parsingResult := handler.ParseEvent(c.Request())
	if event, ok := parsingResult.AsAlarms(); ok {
		return c.JSON(200, map[string]any{
			"status":    "success",
			"eventInfo": fmt.Sprintf("Alarms event received %v", event),
		})
	} else {
		return c.JSON(400, map[string]any{
			"status": "failure",
		})
	}
}
