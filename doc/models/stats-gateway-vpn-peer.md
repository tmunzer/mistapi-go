
# Stats Gateway Vpn Peer

VPN peer path statistics reported by a gateway

## Structure

`StatsGatewayVpnPeer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Whether this VPN peer path is the active redundant path |
| `Jitter` | `*float64` | Optional | Last sampled VPN peer jitter, in milliseconds<br><br>**Constraints**: `>= 0` |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Latency` | `*float64` | Optional | Last sampled VPN peer latency, in milliseconds<br><br>**Constraints**: `>= 0` |
| `Loss` | `*float64` | Optional | Packet loss observed for the VPN peer, as a percentage<br><br>**Constraints**: `>= 0`, `<= 100` |
| `Mos` | `*float64` | Optional | Mean Opinion Score for VPN link quality, from 0 to 5<br><br>**Constraints**: `>= 0`, `<= 5` |
| `Mtu` | `*int` | Optional | Maximum transmission unit for the VPN peer path, in bytes |
| `PeerMac` | `*string` | Optional | Peer router MAC address for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | Peer router interface identifier for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | Peer router name reported for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `PortId` | `*string` | Optional | Local router interface identifier for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | Local router name reported for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `Type` | `*string` | Optional | VPN implementation type for the peer, such as `ipsec` for SRX or `svr` for SSR<br><br>**Constraints**: *Minimum Length*: `1` |
| `Up` | `*bool` | Optional | Whether the VPN peer is currently up |
| `Uptime` | `*int` | Optional | Gateway-reported VPN peer uptime value, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsGatewayVpnPeer := models.StatsGatewayVpnPeer{
        IsActive:             models.ToPointer(false),
        Jitter:               models.ToPointer(float64(43.12)),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Latency:              models.ToPointer(float64(234.58)),
        Loss:                 models.ToPointer(float64(52.32)),
        PeerSiteId:           models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

