
# Ap Search

Access point record returned by device search endpoints

## Structure

`ApSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24Bandwidth` | `*string` | Optional | Current channel bandwidth on the AP 2.4 GHz radio |
| `Band24Channel` | `*int` | Optional | Current channel on the AP 2.4 GHz radio |
| `Band24Power` | `*int` | Optional | Current transmit power on the AP 2.4 GHz radio |
| `Band5Bandwidth` | `*string` | Optional | Current channel bandwidth on the AP 5 GHz radio |
| `Band5Channel` | `*int` | Optional | Current channel on the AP 5 GHz radio |
| `Band5Power` | `*int` | Optional | Current transmit power on the AP 5 GHz radio |
| `Band6Bandwidth` | `*string` | Optional | Current channel bandwidth on the AP 6 GHz radio |
| `Band6Channel` | `*int` | Optional | Current channel on the AP 6 GHz radio |
| `Band6Power` | `*int` | Optional | Current transmit power on the AP 6 GHz radio |
| `Eth0PortSpeed` | `*int` | Optional | Port speed of eth0 |
| `ExtIp` | `*string` | Optional | External IP address observed for AP management traffic |
| `Hostname` | `[]string` | Optional | List of hostnames detected for the AP |
| `InactiveWiredVlans` | `[]int` | Optional | Inactive wired VLAN IDs reported for an AP |
| `Ip` | `*string` | Optional | Management IP address currently assigned to the AP |
| `LastHostname` | `*string` | Optional | Most recent hostname detected for the AP |
| `LldpMgmtAddr` | `*string` | Optional | LLDP management IP address advertised by the upstream neighbor |
| `LldpPortDesc` | `*string` | Optional | LLDP port description advertised by the upstream neighbor |
| `LldpPortId` | `*string` | Optional | LLDP port identifier advertised by the upstream neighbor |
| `LldpPowerAllocated` | `*int` | Optional | Power allocated to the AP by LLDP, in mW |
| `LldpPowerDraw` | `*int` | Optional | Power drawn by the AP as reported through LLDP, in mW |
| `LldpSystemDesc` | `*string` | Optional | LLDP system description advertised by the upstream neighbor |
| `LldpSystemName` | `*string` | Optional | LLDP system name advertised by the upstream neighbor |
| `Mac` | `*string` | Optional | Access point MAC address used to identify the AP in search results |
| `Model` | `*string` | Optional | AP hardware model for this search result |
| `MxedgeId` | `*string` | Optional | Mist Edge id, if AP is connecting to a Mist Edge |
| `MxedgeIds` | `*string` | Optional | Comma separated list of Mist Edge ids, if AP is connecting to a Mist Edge |
| `MxtunnelStatus` | `string` | Required | Current MxTunnel connection status of the AP |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PowerConstrained` | `bool` | Required | Whether the AP is operating with insufficient power |
| `PowerOpmode` | `string` | Required | Operating mode reported when AP power is constrained |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Sku` | `*string` | Optional | Hardware SKU for the AP model |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Uptime` | `*int` | Optional | Device uptime for the AP, in seconds |
| `Version` | `*string` | Optional | Software version currently running on the AP |
| `Wlans` | [`[]models.ApSearchWlan`](../../doc/models/ap-search-wlan.md) | Required | WLAN summaries included in an AP search result |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    apSearch := models.ApSearch{
        Band24Bandwidth:      models.ToPointer("band_24_bandwidth0"),
        Band24Channel:        models.ToPointer(60),
        Band24Power:          models.ToPointer(38),
        Band5Bandwidth:       models.ToPointer("band_5_bandwidth8"),
        Band5Channel:         models.ToPointer(248),
        MxtunnelStatus:       "mxtunnel_status2",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PowerConstrained:     false,
        PowerOpmode:          "power_opmode8",
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Wlans:                []models.ApSearchWlan{
            models.ApSearchWlan{
                Id:                   models.ToPointer(uuid.MustParse("00001c56-0000-0000-0000-000000000000")),
                Ssid:                 models.ToPointer("ssid8"),
            },
        },
    }

}
```

