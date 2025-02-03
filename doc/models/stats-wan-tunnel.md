
# Stats Wan Tunnel

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWanTunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | Authentication algorithm |
| `EncryptAlgo` | `*string` | Optional | Encryption algorithm |
| `IkeVersion` | `*string` | Optional | IKE version |
| `Ip` | `*string` | Optional | IPaddress |
| `LastEvent` | `*string` | Optional | Reason of why the tunnel is down |
| `Mac` | `*string` | Optional | Router mac address |
| `Node` | `*string` | Optional | Node0/node1 |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerHost` | `*string` | Optional | Peer host |
| `PeerIp` | `string` | Required | Peer ip address |
| `Priority` | [`*models.StatsWanTunnelPriorityEnum`](../../doc/models/stats-wan-tunnel-priority-enum.md) | Optional | enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | enum: `gre`, `ipsec` |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
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
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "peer_ip": "peer_ip6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wan_name": "wan",
  "auth_algo": "auth_algo4",
  "encrypt_algo": "encrypt_algo0",
  "ike_version": "ike_version0",
  "ip": "ip0",
  "last_event": "last_event0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

