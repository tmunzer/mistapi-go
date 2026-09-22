
# Stats Mxtunnel

WxLAN or Mist tunnel statistics record

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional, Read-only | MAC address of the AP associated with the tunnel statistics |
| `ForSite` | `*bool` | Optional, Read-only | Whether the tunnel statistics are scoped to a site |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | Firmware update status for a device |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mtu` | `*int` | Optional, Read-only | Maximum transmission unit reported for the tunnel path |
| `MxclusterId` | `*uuid.UUID` | Optional, Read-only | Mist Edge cluster identifier associated with the tunnel statistics |
| `MxedgeId` | `*uuid.UUID` | Optional, Read-only | Mist Edge identifier for the tunnel endpoint |
| `MxtunnelId` | `*uuid.UUID` | Optional, Read-only | Mist tunnel identifier associated with the tunnel statistics |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PeerMxedgeId` | `*uuid.UUID` | Optional, Read-only | MxEdge ID of the peer(mist edge to mist edge tunnel) |
| `RemoteIp` | `string` | Required, Read-only | Remote endpoint IP address for the tunnel |
| `RemotePort` | `*int` | Optional, Read-only | Remote endpoint port for the tunnel |
| `Sessions` | [`[]models.StatsMxtunnelSession`](../../doc/models/stats-mxtunnel-session.md) | Optional, Read-only | Tunnel sessions reported for a WxLAN or Mist tunnel<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `StartTime` | `*int` | Optional, Read-only | Epoch timestamp when the tunnel was established |
| `State` | [`*models.StatsMxtunnelStateEnum`](../../doc/models/stats-mxtunnel-state-enum.md) | Optional, Read-only | enum: `established`, `established_with_sessions`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsMxtunnel := models.StatsMxtunnel{
        Fwupdate:             models.ToPointer(models.FwupdateStat{
        }),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RemoteIp:             "",
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

