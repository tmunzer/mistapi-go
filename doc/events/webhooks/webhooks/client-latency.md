
# Client Latency



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookClientLatency](../../../../doc/models/webhook-client-latency.md).

## Payload Example

```json
{
  "events": [
    {
      "avg_auth": 0.17170219,
      "avg_dhcp": 0.017828934,
      "avg_dns": 0.024532124,
      "max_auth": 0.18170219,
      "max_dhcp": 0.027828934,
      "max_dns": 0.029532124,
      "min_auth": 0.16050219,
      "min_dhcp": 0.015828934,
      "min_dns": 0.022532124,
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1696401600
    }
  ],
  "topic": "client-latency"
}
```

## SDK Usage Example

```go
// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "mistapi/events/webhooks"
)

// WebhooksGinEventHandler receives and processes incoming events
func WebhooksGinEventHandler(c *gin.Context) {
    handler :=  webhooks.NewWebhooksHandler()
    parsingResult := handler.ParseEvent(c.Request)
    if event, ok := parsingResult.AsClientLatency(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("ClientLatency event received %v", event),
        })
    } else if parsingResult.AsUnknown() {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": "Unknown event received",
        })
    } else {
        c.JSON(400, map[string]any{"status": "failure"})
    }
}
```

## Accepted Server Responses

The server should responds with one of the following status codes:

| Status Code | Description |
|  --- | --- |
| 200 | Return a 200 status to indicate that the data was received successfully |

