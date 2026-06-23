
# Ospf Peer Stats Search Results Items

OSPF peer statistic record reported by a router

## Structure

`OspfPeerStatsSearchResultsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeadTime` | `*int` | Optional | Seconds remaining before the neighbor is considered inactive |
| `Mac` | `*string` | Optional | Router MAC address of the device advertising the OSPF peer |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PeerIp` | `*string` | Optional | IP address of the OSPF neighbor |
| `PortId` | `*string` | Optional | Interface on which the OSPF neighbor is learned |
| `Priority` | `*int` | Optional | OSPF priority advertised by the neighbor, from 0 to 255<br><br>**Constraints**: `>= 0`, `<= 255` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `State` | `*string` | Optional | Eg. full, down, 2way, init, exstart, exchange, loading |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Up` | `*bool` | Optional | True if state is full |
| `VrfName` | `*string` | Optional | Instance name, e.g. master |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ospfPeerStatsSearchResultsItems := models.OspfPeerStatsSearchResultsItems{
        DeadTime:             models.ToPointer(96),
        Mac:                  models.ToPointer("mac8"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PeerIp:               models.ToPointer("peer_ip6"),
        PortId:               models.ToPointer("port_id4"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

