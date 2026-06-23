
# Stats Wan Tunnel

WAN tunnel statistics record returned by tunnel search

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | Authentication algorithm negotiated for the tunnel |
| `EncryptAlgo` | `*string` | Optional | Encryption algorithm negotiated for the tunnel |
| `IkeVersion` | `*string` | Optional | IKE version used to establish the tunnel |
| `Ip` | `*string` | Optional | Local IP address used by the tunnel |
| `LastEvent` | `*string` | Optional | Most recent reason the tunnel went down |
| `Mac` | `*string` | Optional | Router MAC address reporting the tunnel statistics |
| `Node` | `*string` | Optional | HA node handling the tunnel, such as node0 or node1 |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PeerHost` | `*string` | Optional | Hostname or configured peer host of the remote tunnel endpoint |
| `PeerIp` | `string` | Required | IP address of the remote tunnel endpoint |
| `Priority` | [`*models.TunnelPriorityEnum`](../../doc/models/tunnel-priority-enum.md) | Optional | Relative preference assigned to the tunnel. enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | Tunnel protocol used for the connection. enum: `gre`, `ipsec` |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TunnelName` | `*string` | Optional | Name of the Mist-managed WAN tunnel |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | Indicates whether the tunnel is currently up |
| `Uptime` | `*int` | Optional | Duration since the tunnel security association was established |
| `WanName` | `*string` | Optional | Name of the WAN interface carrying the tunnel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWanTunnel := models.StatsWanTunnel{
        AuthAlgo:             models.ToPointer("auth_algo8"),
        EncryptAlgo:          models.ToPointer("encrypt_algo4"),
        IkeVersion:           models.ToPointer("ike_version4"),
        Ip:                   models.ToPointer("ip4"),
        LastEvent:            models.ToPointer("last_event4"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PeerIp:               "peer_ip0",
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        WanName:              models.ToPointer("wan"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

