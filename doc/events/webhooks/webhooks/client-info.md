
# Client Info



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookClientInfo](../../../../doc/models/webhook-client-info.md).

## Payload Example

```json
{
  "events": [
    {
      "hostname": "service.company.net",
      "ip": "21.0.128.151",
      "mac": "6ebaa47a3fd4",
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "site_id": "6e77a2ea-d579-471c-9056-5ff5b4ed70ed",
      "timestamp": 1703003296
    }
  ],
  "topic": "client-info"
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
    if event, ok := parsingResult.AsClientInfo(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("ClientInfo event received %v", event),
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

