
# Webhook Alarm Event

Alarm event delivered inside an `alarms` webhook payload

## Structure

`WebhookAlarmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | `[]string` | Optional, Read-only | Representative AP MAC address samples included for an alarm webhook event |
| `Bssids` | `[]string` | Optional, Read-only | Representative BSSID samples included for an alarm webhook event |
| `Count` | `*int` | Optional, Read-only | If present, number of events of this type that occurred in the current interval; defaults to 1 |
| `EventId` | `*uuid.UUID` | Optional, Read-only | Unique identifier for this alarm webhook event |
| `ForSite` | `*bool` | Optional, Read-only | Whether this alarm event is scoped to a site rather than only to the organization |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Ssids` | `[]string` | Optional, Read-only | Representative wireless SSID samples included for an alarm webhook event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Read-only | Alarm type key for this event |
| `Update` | `*bool` | Optional, Read-only | Whether this payload updates an alarm event that was sent earlier; defaults to false |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookAlarmEvent := models.WebhookAlarmEvent{
        Aps:                  []string{
            "aps5",
        },
        Bssids:               []string{
            "bssids8",
            "bssids7",
            "bssids6",
        },
        Count:                models.ToPointer(142),
        EventId:              models.ToPointer(uuid.MustParse("00001ed2-0000-0000-0000-000000000000")),
        ForSite:              models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:            float64(211.12),
        Type:                 "type4",
    }

}
```

