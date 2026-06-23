
# Stats Gateway Wan Tunnel

WAN tunnel statistics reported by a gateway

## Structure

`StatsGatewayWanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | Authentication algorithm negotiated for the tunnel |
| `EncryptAlgo` | `*string` | Optional | Encryption algorithm negotiated for the tunnel |
| `IkeVersion` | `*string` | Optional | IKE version used to establish the tunnel |
| `Ip` | `*string` | Optional | Local IP address used by the tunnel |
| `LastEvent` | `*string` | Optional | Most recent reason the tunnel went down |
| `LastFlapped` | `*float64` | Optional | Indicates when the tunnel last flapped |
| `Node` | `*string` | Optional | HA node handling the tunnel, such as node0 or node1 |
| `PeerHost` | `*string` | Optional | Hostname or configured peer host of the remote tunnel endpoint |
| `PeerIp` | `*string` | Optional | IP address of the remote tunnel endpoint |
| `Priority` | [`*models.TunnelPriorityEnum`](../../doc/models/tunnel-priority-enum.md) | Optional | Relative preference assigned to the tunnel. enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | Tunnel protocol used for the connection. enum: `gre`, `ipsec` |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `TunnelName` | `*string` | Optional | Name of the Mist-managed WAN tunnel |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | Indicates whether the tunnel is currently up |
| `Uptime` | `*int` | Optional | Duration since the tunnel security association was established |
| `WanName` | `*string` | Optional | Name of the WAN interface carrying the tunnel |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayWanTunnel := models.StatsGatewayWanTunnel{
        AuthAlgo:             models.ToPointer("auth_algo4"),
        EncryptAlgo:          models.ToPointer("encrypt_algo2"),
        IkeVersion:           models.ToPointer("ike_version2"),
        Ip:                   models.ToPointer("ip2"),
        LastEvent:            models.ToPointer("last_event2"),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        WanName:              models.ToPointer("wan"),
    }

}
```

