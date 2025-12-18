
# Audits



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookAudits](../../../../doc/models/webhook-audits.md).

## Payload Example

```json
{
  "events": [
    {
      "admin_name": "admin_name8",
      "device_id": "00000380-0000-0000-0000-000000000000",
      "id": "0000122a-0000-0000-0000-000000000000",
      "message": "message0",
      "org_id": "00001302-0000-0000-0000-000000000000",
      "site_id": "00000290-0000-0000-0000-000000000000",
      "src_ip": "src_ip6",
      "timestamp": 157.68,
      "admin_id": "0000104e-0000-0000-0000-000000000000",
      "after": {
        "key1": "val1",
        "key2": "val2"
      },
      "before": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "audits"
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
    if event, ok := parsingResult.AsAudits(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Audits event received %v", event),
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

