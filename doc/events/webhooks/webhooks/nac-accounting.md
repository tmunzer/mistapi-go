
# Nac Accounting



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookNacAccounting](../../../../doc/models/webhook-nac-accounting.md).

## Payload Example

```json
{
  "events": [
    {
      "ap": "5c5b355005be",
      "auth_type": "eap-tls",
      "bssid": "5c5b35546bb4",
      "client_ip": "172.16.87.4",
      "client_type": "wireless",
      "mac": "6e795836d5f9",
      "nas_vendor": "juniper-mist",
      "org_id": "625aba64-fe72-4b14-8985-cbf31ec3d78a",
      "rx_pkts": 9523,
      "site_id": "ec9d1e85-af24-43f9-8d65-d620580e8631",
      "ssid": "Test-CMR SSID",
      "timestamp": 1547235620.89,
      "tx_pkts": 5233,
      "type": "NAC_ACCOUNTING_STOP",
      "username": "hi"
    }
  ],
  "topic": "nac-accounting"
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
    if event, ok := parsingResult.AsNacAccounting(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("NacAccounting event received %v", event),
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

