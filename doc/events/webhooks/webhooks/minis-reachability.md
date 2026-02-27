
# Minis Reachability



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookMinisReachability](../../../../doc/models/webhook-minis-reachability.md).

## Payload Example

```json
{
  "events": [
    {
      "avg_latency": 12.5,
      "device_mac": "7cb68d8f0440",
      "loss_percentage": 0.0,
      "max_latency": 15.2,
      "min_latency": 10.1,
      "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4",
      "probe_name": "google ping",
      "probe_target": "google.com",
      "probe_type": "reachability",
      "protocol": "icmp",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "test_type": "ping",
      "timestamp": 1547235620.89,
      "vlan": 12
    }
  ],
  "topic": "minis-reachability"
}
```

## SDK Usage Example

```go
// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "mistapi/events/webhooks"
)

// WebhooksGinEventHandler receives and processes incoming events
func WebhooksGinEventHandler(c *gin.Context) {
    handler :=  webhooks.NewWebhooksHandler()
    parsingResult := handler.ParseEvent(c.Request)
    if event, ok := parsingResult.AsMinisReachability(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("MinisReachability event received %v", event),
        })
    } else if parsingResult.AsUnknown() {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": "Unknown event received",
        })
    } else {
        c.JSON(400, map[string]any{
            "status": "failure",
        })
    }
}
```

## Accepted Server Responses

The server should responds with one of the following status codes:

| Status Code | Description |
|  --- | --- |
| 200 | Return a 200 status to indicate that the data was received successfully |

