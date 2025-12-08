
# Mxedge Events



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookMxedgeEvents](../../../../doc/models/webhook-mxedge-events.md).

## Payload Example

```json
{
  "events": [
    {
      "audit_id": "8912e5cb-8ddd-41f7-be5f-476a7abbf658",
      "component": null,
      "mxedge_id": "00000000-0000-0000-1000-020000230522",
      "mxedge_name": "demo123",
      "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4",
      "timestamp": "1763546876.209649",
      "type": "ME_CONFIG_CHANGED_BY_USER",
      "device_type": "device_type0",
      "from_version": "from_version2",
      "mac": "mac4"
    },
    {
      "audit_id": "48efa5bf-d290-4e93-80ca-4dbf72f4187a",
      "component": null,
      "mxedge_id": "00000000-0000-0000-1000-020000a5fca1",
      "mxedge_name": "test123",
      "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4",
      "timestamp": "1763546876.417778",
      "type": "ME_CONFIG_CHANGED_BY_USER",
      "device_type": "device_type0",
      "from_version": "from_version2",
      "mac": "mac4"
    }
  ],
  "topic": "mxedge-events"
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
    if event, ok := parsingResult.AsMxedgeEvents(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("MxedgeEvents event received %v", event),
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

