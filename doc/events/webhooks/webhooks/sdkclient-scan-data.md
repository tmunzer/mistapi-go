
# Sdkclient Scan Data



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookSdkclientScanData](../../../../doc/models/webhook-sdkclient-scan-data.md).

## Payload Example

```json
{
  "events": [
    {
      "connection_ap": "5c5b352f587e",
      "connection_band": "2.4",
      "connection_bssid": "5c5b352b51b4",
      "connection_channel": 11,
      "connection_rssi": -87.0,
      "last_seen": 1592333828.0,
      "mac": "70ef0071535f",
      "scan_data": [
        {
          "ap": "5c5b352f587e",
          "band": "2.4",
          "bssid": "5c5b352b51b4",
          "channel": 11,
          "rssi": -87.0,
          "ssid": "mist-wifi",
          "timestamp": 1592333828.0
        },
        {
          "ap": "5c5b352f587e",
          "band": "5",
          "bssid": "5c5b352b51b8",
          "channel": 36,
          "rssi": -75.0,
          "ssid": "mist-wifi",
          "timestamp": 1592333828.0
        }
      ],
      "site_id": "93986f10-773b-42be-9438-8d3e6d127f1a"
    }
  ],
  "topic": "sdkclient-scan-data"
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
    if event, ok := parsingResult.AsSdkclientScanData(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("SdkclientScanData event received %v", event),
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

