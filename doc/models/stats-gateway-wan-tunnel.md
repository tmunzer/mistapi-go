
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
| `Ip` | `*string` | Optional | IPaddress |
| `LastEvent` | `*string` | Optional | Reason of why the tunnel is down |
| `LastFlapped` | `*float64` | Optional | Indicates when the port was last flapped |
| `Node` | `*string` | Optional | Node0/node1 |
| `PeerHost` | `*string` | Optional | Peer host |
| `PeerIp` | `*string` | Optional | Peer ip address |
| `Priority` | [`*models.StatsWanTunnelPriorityEnum`](../../doc/models/stats-wan-tunnel-priority-enum.md) | Optional | enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | enum: `gre`, `ipsec` |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `TunnelName` | `*string` | Optional | Mist Tunnel Name |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | Duration from first (or last) SA was established |
| `WanName` | `*string` | Optional | WAN interface name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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

