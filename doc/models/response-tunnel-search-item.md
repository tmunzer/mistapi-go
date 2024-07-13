
# Response Tunnel Search Item

## Structure

`ResponseTunnelSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `LastSeen` | `*float64` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `MxedgeId` | `*uuid.UUID` | Optional | - |
| `MxtunnelId` | `*uuid.UUID` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerMxedgeId` | `*uuid.UUID` | Optional | MxEdge ID of the peer(mist edge to mist edge tunnel) |
| `RemoteIp` | `*string` | Optional | - |
| `RemotePort` | `*int` | Optional | - |
| `RxControlPkts` | `*int` | Optional | - |
| `Sessions` | [`[]models.MxtunnelStatsSession`](../../doc/models/mxtunnel-stats-session.md) | Optional | list of sessions<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `State` | [`*models.MxtunnelStatsStateEnum`](../../doc/models/mxtunnel-stats-state-enum.md) | Optional | - |
| `TxControlPkts` | `*int` | Optional | - |
| `Uptime` | `*int` | Optional | duration from first (or last) SA was established |
| `AuthAlgo` | `*string` | Optional | authentication algorithm |
| `EncryptAlgo` | `*string` | Optional | encryption algorithm |
| `IkeVersion` | `*string` | Optional | ike version |
| `Ip` | `*string` | Optional | ip address |
| `LastEvent` | `*string` | Optional | reason of why the tunnel is down |
| `Mac` | `*string` | Optional | router mac address |
| `Node` | `*string` | Optional | node0/node1 |
| `PeerHost` | `*string` | Optional | peer host |
| `PeerIp` | `*string` | Optional | peer ip address |
| `Priority` | [`*models.WanTunnelStatsPriorityEnum`](../../doc/models/wan-tunnel-stats-priority-enum.md) | Optional | - |
| `Protocol` | [`*models.WanTunnelProtocolEnum`](../../doc/models/wan-tunnel-protocol-enum.md) | Optional | - |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `TunnelName` | `*string` | Optional | Mist Tunnel Name |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `WanName` | `*string` | Optional | wan interface name |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wan_name": "wan",
  "ap": "ap2",
  "for_site": false,
  "last_seen": 240.02,
  "mxcluster_id": "000012b8-0000-0000-0000-000000000000",
  "mxedge_id": "000011ec-0000-0000-0000-000000000000"
}
```

