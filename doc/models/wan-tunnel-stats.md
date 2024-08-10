
# Wan Tunnel Stats

## Structure

`WanTunnelStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | `*string` | Optional | authentication algorithm |
| `EncryptAlgo` | `*string` | Optional | encryption algorithm |
| `IkeVersion` | `*string` | Optional | ike version |
| `Ip` | `*string` | Optional | ip address |
| `LastEvent` | `*string` | Optional | reason of why the tunnel is down |
| `Mac` | `*string` | Optional | router mac address |
| `Node` | `*string` | Optional | node0/node1 |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerHost` | `*string` | Optional | peer host |
| `PeerIp` | `*string` | Optional | peer ip address |
| `Priority` | [`*models.WanTunnelStatsPriorityEnum`](../../doc/models/wan-tunnel-stats-priority-enum.md) | Optional | enum: `primary`, `secondary` |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | enum: `gre`, `ipsec` |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TunnelName` | `*string` | Optional | Mist Tunnel Name |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | duration from first (or last) SA was established |
| `WanName` | `*string` | Optional | wan interface name |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wan_name": "wan",
  "auth_algo": "auth_algo8",
  "encrypt_algo": "encrypt_algo8",
  "ike_version": "ike_version8",
  "ip": "ip8",
  "last_event": "last_event8"
}
```

