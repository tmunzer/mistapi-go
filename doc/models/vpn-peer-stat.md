
# Vpn Peer Stat

VPN peer path statistics returned by organization VPN peer searches

## Structure

`VpnPeerStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IsActive` | `*bool` | Optional | Whether this VPN peer path is the active redundant path |
| `Jitter` | `*float64` | Optional | Last sampled VPN peer jitter, in milliseconds<br><br>**Constraints**: `>= 0` |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Latency` | `*float64` | Optional | Last sampled VPN peer latency, in milliseconds<br><br>**Constraints**: `>= 0` |
| `Loss` | `*float64` | Optional | Packet loss in percentage<br><br>**Constraints**: `>= 0`, `<= 100` |
| `Mac` | `*string` | Optional | Local router MAC address for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `Mos` | `*float64` | Optional | Mean Opinion Score, a measure of the quality of the VPN link<br><br>**Constraints**: `>= 0`, `<= 5` |
| `Mtu` | `*int` | Optional | Maximum transmission unit for the VPN peer path, in bytes |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PeerMac` | `*string` | Optional | Peer router MAC address for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerPortId` | `*string` | Optional | Peer router interface identifier for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerRouterName` | `*string` | Optional | Peer router name reported for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `PeerSiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `PortId` | `*string` | Optional | Local router interface identifier for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `RouterName` | `*string` | Optional | Local router name reported for the VPN link<br><br>**Constraints**: *Minimum Length*: `1` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
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
    vpnPeerStat := models.VpnPeerStat{
        IsActive:             models.ToPointer(false),
        Jitter:               models.ToPointer(float64(83.56)),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Latency:              models.ToPointer(float64(194.14)),
        Loss:                 models.ToPointer(float64(92.76)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PeerSiteId:           models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

