
# Device Updowns



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookDeviceUpdowns](../../../../doc/models/webhook-device-updowns.md).

## Payload Example

```json
{
  "events": [
    {
      "ap": "string",
      "ap_name": "string",
      "for_site": true,
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
      "site_name": "string",
      "timestamp": 0.0,
      "type": "string"
    }
  ],
  "topic": "device-updowns"
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
    if event, ok := parsingResult.AsDeviceUpdowns(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("DeviceUpdowns event received %v", event),
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

