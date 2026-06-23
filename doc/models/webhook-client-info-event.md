
# Webhook Client Info Event

Client identity and addressing details delivered by a `client-info` webhook

## Structure

`WebhookClientInfoEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | Client hostname reported in the event |
| `Ip` | `*string` | Optional | Client IP address reported in the event |
| `Mac` | `*string` | Optional | Client MAC address reported in the event |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientInfoEvent := models.WebhookClientInfoEvent{
        Hostname:             models.ToPointer("service.company.net"),
        Ip:                   models.ToPointer("ip6"),
        Mac:                  models.ToPointer("mac6"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

