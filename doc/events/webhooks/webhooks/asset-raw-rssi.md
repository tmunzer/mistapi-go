
# Asset Raw Rssi



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookAssetRawRssi](../../../../doc/models/webhook-asset-raw-rssi.md).

## Payload Example

```json
{
  "events": [
    {
      "ap_loc": [
        36.03303862386182,
        43.57022468463291,
        2.75
      ],
      "beam": 8,
      "device_id": "00000000-0000-0000-1000-ac2316eca70b",
      "ibeacon_major": 2,
      "ibeacon_minor": 2121,
      "ibeacon_uuid": "ac950d7b-af31-42d2-be7c-e15639fab2cd",
      "is_asset": true,
      "mac": "ed2cc53f2770",
      "map_id": "bd42f0c3-2e6a-4f8a-ac2d-d34e268c1418",
      "org_id": "9c3e516c-397d-11e6-ae35-0242ac110008",
      "rssi": -79,
      "service_packets": [
        {
          "service_data": "010441060606fe3d35700601cecbd902512f000001",
          "service_uuid": "UUID"
        }
      ],
      "site_id": "27ea2f07-6fe6-4eab-be1b-b8e3ce083d67",
      "timestamp": 1661300746
    },
    {
      "ap_loc": [
        36.03303862386182,
        43.57022468463291,
        2.75
      ],
      "beam": 7,
      "device_id": "00000000-0000-0000-1000-ac2316eca70b",
      "is_asset": true,
      "mac": "ed2cc53f2771",
      "map_id": "bd42f0c3-2e6a-4f8a-ac2d-d34e268c1418",
      "mfg_company_id": 243,
      "mfg_data": "00000000000000000000000000",
      "org_id": "9c3e516c-397d-11e6-ae35-0242ac110008",
      "rssi": -74,
      "site_id": "27ea2f07-6fe6-4eab-be1b-b8e3ce083d67",
      "timestamp": 1661300746,
      "ibeacon_major": 178,
      "ibeacon_minor": 40
    }
  ],
  "topic": "asset-raw-rssi"
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
    if event, ok := parsingResult.AsAssetRawRssi(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("AssetRawRssi event received %v", event),
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

