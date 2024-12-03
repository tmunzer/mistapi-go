
# Optional Stat Wan Tunnel

*This model accepts additional fields of type interface{}.*

## Structure

`OptionalStatWanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | authentication algorithm |
| `EncryptAlgo` | `*string` | Optional | encryption algorithm |
| `IkeVersion` | `*string` | Optional | ike version |
| `Ip` | `*string` | Optional | ip address |
| `LastEvent` | `*string` | Optional | reason of why the tunnel is down |
| `LastFlapped` | `*float64` | Optional | indicates when the port was last flapped |
| `Node` | `*string` | Optional | node0/node1 |
| `PeerHost` | `*string` | Optional | peer host |
| `PeerIp` | `*string` | Optional | peer ip address |
| `Priority` | [`*models.StatsWanTunnelPriorityEnum`](../../doc/models/stats-wan-tunnel-priority-enum.md) | Optional | enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | enum: `gre`, `ipsec` |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `TunnelName` | `*string` | Optional | Mist Tunnel Name |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | duration from first (or last) SA was established |
| `WanName` | `*string` | Optional | wan interface name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "wan_name": "wan",
  "auth_algo": "auth_algo8",
  "encrypt_algo": "encrypt_algo8",
  "ike_version": "ike_version8",
  "ip": "ip8",
  "last_event": "last_event8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

