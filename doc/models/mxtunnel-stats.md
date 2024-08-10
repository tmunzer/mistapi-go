
# Mxtunnel Stats

MxTunnels statistics

## Structure

`MxtunnelStats`

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
| `State` | [`*models.MxtunnelStatsStateEnum`](../../doc/models/mxtunnel-stats-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `TxControlPkts` | `*int` | Optional | - |
| `Uptime` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap": "ap4",
  "for_site": false,
  "last_seen": 153.2,
  "mxcluster_id": "000017de-0000-0000-0000-000000000000",
  "mxedge_id": "00001712-0000-0000-0000-000000000000"
}
```

