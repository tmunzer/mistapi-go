
# Zone



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookZone](../../../../doc/models/webhook-zone.md).

## Payload Example

```json
{
  "events": [
    {
      "id": "487f6eca-6276-4993-bfeb-d3cbbbba2f08",
      "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
      "name": "string",
      "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
      "timestamp": 0.0,
      "type": "string",
      "asset_id": "00001e56-0000-0000-0000-000000000000",
      "mac": "mac4",
      "trigger": "enter",
      "zone_id": "00000f60-0000-0000-0000-000000000000"
    }
  ],
  "topic": "zone"
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
    if event, ok := parsingResult.AsZone(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Zone event received %v", event),
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

