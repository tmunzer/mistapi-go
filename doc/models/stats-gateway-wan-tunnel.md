
# Stats Gateway Wan Tunnel

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGatewayWanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | Authentication algorithm |
| `EncryptAlgo` | `*string` | Optional | Encryption algorithm |
| `IkeVersion` | `*string` | Optional | IKE version |
| `Ip` | `*string` | Optional | IP Address |
| `LastEvent` | `*string` | Optional | Reason of why the tunnel is down |
| `LastFlapped` | `*float64` | Optional | Indicates when the port was last flapped |
| `Node` | `*string` | Optional | Node0/node1 |
| `PeerHost` | `*string` | Optional | Peer host |
| `PeerIp` | `*string` | Optional | Peer ip address |
| `Priority` | [`*models.StatsWanTunnelPriorityEnum`](../../doc/models/stats-wan-tunnel-priority-enum.md) | Optional | enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | enum: `gre`, `ipsec` |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `TunnelName` | `*string` | Optional | Mist Tunnel Name |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | Duration from first (or last) SA was established |
| `WanName` | `*string` | Optional | WAN interface name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "wan_name": "wan",
  "auth_algo": "auth_algo8",
  "encrypt_algo": "encrypt_algo4",
  "ike_version": "ike_version4",
  "ip": "ip4",
  "last_event": "last_event4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

