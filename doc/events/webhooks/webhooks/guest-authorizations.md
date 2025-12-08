
# Guest Authorizations



## Headers

This event's request contains the following headers.

| Name |
|  --- |
| X-Mist-Signature-v2 |
| X-Mist-Signature |
| Content-Type |

## Payload Type

This event's request payload is of type [*models.WebhookGuestAuthorizations](../../../../doc/models/webhook-guest-authorizations.md).

## Payload Example

```json
{
  "events": [
    {
      "ap": "5c5b350e55c8",
      "auth_method": "passphrase",
      "authorized_expiring_time": 1677076639,
      "authorized_time": 1677076519,
      "carrier": "docomo",
      "client": "ac2316eca70a",
      "company": "MIST",
      "email": "abcd@abcd.com",
      "field1": "field1 value",
      "field2": "field2 value",
      "field3": "field3 value",
      "field4": "field4 value",
      "mobile": "+0123456789",
      "name": "Dr Strange",
      "org_id": "1688605f-916a-47a1-8c68-f19618300a08",
      "site_id": "ec3b5624-73f1-4ed3-b3fd-5ba3ee40368a",
      "sms_gateway": "Telstra",
      "sponsor_email": "sponsor@gmail.com",
      "ssid": "Portal Auth",
      "wlan_id": "7681be9a-044a-4622-90cf-3accde5ad853"
    }
  ],
  "topic": "guest-authorizations"
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
    if event, ok := parsingResult.AsGuestAuthorizations(); ok {
        c.JSON(200, map[string]any{
            "status":    "success",
            "eventInfo": fmt.Sprintf("GuestAuthorizations event received %v", event),
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

