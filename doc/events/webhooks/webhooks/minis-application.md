
# Minis Application



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookMinisApplication](../../../../doc/models/webhook-minis-application.md).

## Payload Example

```json
{
  "events": [
    {
      "device_mac": "003e7316fd96",
      "ip": "172.217.22.195",
      "latency": 74,
      "org_id": "8aa21779-1178-4357-b3e0-42c02b93b870",
      "probe_name": "connectivitycheck.gstatic.com",
      "probe_type": "application",
      "site_id": "2bf12442-1558-41bd-849e-738d6d4aa1a3",
      "src_ip": "192.168.1.136",
      "success": false,
      "test_type": "curl",
      "timestamp": 1775641319,
      "vlan": 1
    }
  ],
  "topic": "minis-application"
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
    if event, ok := parsingResult.AsMinisApplication(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("MinisApplication event received %v", event),
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

