
# Mxtunnel

Mist Tunnel configuration for tunneling AP user VLANs to Mist Edge clusters

*This model accepts additional fields of type interface{}.*

## Structure

`Mxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnchorMxtunnelIds` | `[]uuid.UUID` | Optional | List of anchor mxtunnels used for forming edge to edge tunnels |
| `AutoPreemption` | [`*models.AutoPreemption`](../../doc/models/auto-preemption.md) | Optional | Schedule to preempt AP tunnels that are not connected to their preferred peer |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this Mist Tunnel is scoped to a site |
| `HelloInterval` | `models.Optional[int]` | Optional | In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by `hello_retries`.<br><br>**Default**: `60`<br><br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `models.Optional[int]` | Optional | Number of missed hello heartbeats before an AP tries another tunnel peer<br><br>**Default**: `7`<br><br>**Constraints**: `>= 2`, `<= 30` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Ipsec` | [`*models.MxtunnelIpsec`](../../doc/models/mxtunnel-ipsec.md) | Optional | IPsec settings for a Mist Tunnel |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 1500` |
| `MxclusterIds` | `[]uuid.UUID` | Optional | Mist Edge cluster IDs that host this Mist Tunnel |
| `Name` | `models.Optional[string]` | Optional | Display name of the Mist Tunnel |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Protocol` | [`*models.MxtunnelProtocolEnum`](../../doc/models/mxtunnel-protocol-enum.md) | Optional | Encapsulation protocol used for Mist Tunnel traffic. enum: `ip`, `udp`<br><br>**Default**: `"udp"` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `VlanIds` | `[]int` | Optional | List of VLAN IDs carried by this Mist Tunnel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxtunnel := models.Mxtunnel{
        AnchorMxtunnelIds:    []uuid.UUID{
            uuid.MustParse("00000338-0000-0000-0000-000000000000"),
            uuid.MustParse("00000339-0000-0000-0000-000000000000"),
        },
        AutoPreemption:       models.ToPointer(models.AutoPreemption{
            DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_MON),
            Enabled:              models.ToPointer(false),
            TimeOfDay:            models.ToPointer("time_of_day4"),
        }),
        CreatedTime:          models.ToPointer(float64(69.9)),
        ForSite:              models.ToPointer(false),
        HelloInterval:        models.NewOptional(models.ToPointer(60)),
        HelloRetries:         models.NewOptional(models.ToPointer(7)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mtu:                  models.ToPointer(0),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Protocol:             models.ToPointer(models.MxtunnelProtocolEnum_UDP),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

