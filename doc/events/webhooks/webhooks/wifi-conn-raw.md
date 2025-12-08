
# Wifi Conn Raw



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookWifiConnRaw](../../../../doc/models/webhook-wifi-conn-raw.md).

## Payload Example

```json
{
  "events": [
    {
      "ap_id": "ac-23-16-ec-a7-0b",
      "ap_loc": [
        36.03303862386182,
        43.57022468463291,
        2.75
      ],
      "client_id": "28-f0-76-2d-22-1e",
      "connected_site": false,
      "mapid": "bd42f0c3-2e6a-4f8a-ac2d-d34e268c1418",
      "orgid": "9c3e516c-397d-11e6-ae35-0242ac110008",
      "packets": [
        {
          "band": "5GHz",
          "rssi": -92
        }
      ],
      "siteid": "27ea2f07-6fe6-4eab-be1b-b8e3ce083d67",
      "extended_info_list": [
        {
          "frame_ctrl": 248,
          "payload": "payload2",
          "sequence_ctrl": 42
        },
        {
          "frame_ctrl": 248,
          "payload": "payload2",
          "sequence_ctrl": 42
        }
      ]
    },
    {
      "ap_id": "ac-23-16-ec-a7-0b",
      "ap_loc": [
        36.03303862386182,
        43.57022468463291,
        2.75
      ],
      "client_id": "38-f9-d3-99-08-6e",
      "connected_site": false,
      "extended_info_list": [
        {
          "frame_ctrl": 776,
          "payload": "010441060606fe3d35700601cecbd902512f000001",
          "sequence_ctrl": 8432
        }
      ],
      "mapid": "bd42f0c3-2e6a-4f8a-ac2d-d34e268c1418",
      "orgid": "9c3e516c-397d-11e6-ae35-0242ac110008",
      "packets": [
        {
          "band": "5GHz",
          "rssi": -94.333336
        }
      ],
      "siteid": "27ea2f07-6fe6-4eab-be1b-b8e3ce083d67"
    }
  ],
  "topic": "wifi-conn-raw"
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
    if event, ok := parsingResult.AsWifiConnRaw(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("WifiConnRaw event received %v", event),
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

