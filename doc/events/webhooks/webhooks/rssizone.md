
# Rssizone



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookRssizone](../../../../doc/models/webhook-rssizone.md).

## Payload Example

```json
{
  "events": [
    {
      "mac": "500291xxxxxx",
      "map_id": "f5d26c7f-1670-4921-xxxx-xxxxxxxxxxxx",
      "rssizone_id": "e38f8e76-40db-4144-xxxx-xxxxxxxxxxxx",
      "site_id": "f5fcbee5-fbca-45b3-xxxx-xxxxxxxxxxxx",
      "timestamp": 1694158990.986472,
      "trigger": "enter",
      "type": "wifi"
    }
  ],
  "topic": "rssizone"
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
    if event, ok := parsingResult.AsRssizone(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Rssizone event received %v", event),
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

