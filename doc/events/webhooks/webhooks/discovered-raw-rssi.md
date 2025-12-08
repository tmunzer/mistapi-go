
# Discovered Raw Rssi



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookDiscoveredRawRssi](../../../../doc/models/webhook-discovered-raw-rssi.md).

## Payload Example

```json
{
  "events": [
    {
      "ap_loc": [
        0.0
      ],
      "beam": 0,
      "device_id": "3bafab7b-4400-4bcf-8e6e-09f954699940",
      "ibeacon_major": 1,
      "ibeacon_minor": 1,
      "ibeacon_uuid": "1f89bc00-d0af-481b-82fe-a6629259a39f",
      "is_asset": true,
      "mac": "string",
      "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
      "mfg_company_id": "string",
      "mfg_data": "string",
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "rssi": 0.0,
      "service_packets": [
        {
          "service_data": "string",
          "service_uuid": "7138cc00-c611-4dec-a05e-5c4b1cae13c0"
        }
      ],
      "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
      "timestamp": 0
    }
  ],
  "topic": "string"
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
    if event, ok := parsingResult.AsDiscoveredRawRssi(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("DiscoveredRawRssi event received %v", event),
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

