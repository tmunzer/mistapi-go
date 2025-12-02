
# Nac Events



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookNacEvents](../../../../doc/models/webhook-nac-events.md).

## Payload Example

```json
{
  "events": [
    {
      "ap": "5c5b35513227",
      "auth_type": "eap-teap",
      "bssid": "5c5b355fafcc",
      "dryrun_nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264012",
      "dryrun_nacrule_matched": true,
      "idp_id": "912ef72e-2239-4996-b81e-469e87a27cd6",
      "idp_role": [
        "itsuperusers",
        "vip"
      ],
      "mac": "ac3eb179e535",
      "nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62",
      "nacrule_matched": true,
      "nas_vendor": "juniper-mist",
      "org_id": "27547ac2-d114-4e04-beb1-f3f1e6e81ec6",
      "random_mac": "true",
      "resp_attrs": [
        "Tunnel-Type=VLAN",
        "Tunnel-Medium-Type=IEEE-802",
        "Tunnel-Private-Group-Id=750",
        "User-Name=anonymous"
      ],
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ssid": "##mist_nac",
      "timestamp": 1691512031.358188,
      "type": "NAC_CLIENT_PERMIT",
      "username": "user@deaflyz.net",
      "vlan": "750",
      "client_type": "wired",
      "device_mac": "device_mac4"
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
    if event, ok := parsingResult.AsNacEvents(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("NacEvents event received %v", event),
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

