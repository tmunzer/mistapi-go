
# Location



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookLocation](../../../../doc/models/webhook-location.md).

## Payload Example

```json
{
  "events": [
    {
      "battery_voltage": 0,
      "eddystone_uid_instance": "string",
      "eddystone_uid_namespace": "string",
      "eddystone_url_url": "string",
      "ibeacon_major": 1,
      "ibeacon_minor": 1,
      "ibeacon_uuid": "1f89bc00-d0af-481b-82fe-a6629259a39f",
      "id": "487f6eca-6276-4993-bfeb-e3cbbbba3f08",
      "mac": "string",
      "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
      "mfg_company_id": 0,
      "mfg_data": "string",
      "name": "string",
      "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
      "timestamp": 0.0,
      "type": "string",
      "x": 0.0,
      "y": 0.0
    }
  ],
  "topic": "location"
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
    if event, ok := parsingResult.AsLocation(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("Location event received %v", event),
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

