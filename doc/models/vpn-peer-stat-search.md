
# Vpn Peer Stat Search

Search response containing VPN peer path statistics

## Structure

`VpnPeerStatSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Search window end timestamp, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of VPN peer statistics returned |
| `Next` | `*string` | Optional | URL for the next page of VPN peer statistics, when more results are available |
| `Results` | [`[]models.VpnPeerStat`](../../doc/models/vpn-peer-stat.md) | Required | VPN peer statistic rows returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Search window start timestamp, in epoch seconds |
| `Total` | `int` | Required | Number of VPN peer statistics matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    vpnPeerStatSearch := models.VpnPeerStatSearch{
        End:                  float64(117.08),
        Limit:                238,
        Next:                 models.ToPointer("next8"),
        Results:              []models.VpnPeerStat{
            models.VpnPeerStat{
                IsActive:             models.ToPointer(false),
                Jitter:               models.ToPointer(float64(27.56)),
                LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
                Latency:              models.ToPointer(float64(250.14)),
                Loss:                 models.ToPointer(float64(36.76)),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                PeerSiteId:           models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Start:                float64(73.14),
        Total:                76,
    }

}
```

