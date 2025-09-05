
# Stats Mxtunnel

MxTunnels statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mtu` | `*int` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `MxedgeId` | `*uuid.UUID` | Optional | - |
| `MxtunnelId` | `*uuid.UUID` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PeerMxedgeId` | `*uuid.UUID` | Optional | MxEdge ID of the peer(mist edge to mist edge tunnel) |
| `RemoteIp` | `string` | Required | - |
| `RemotePort` | `*int` | Optional | - |
| `RxControlPkts` | `*int` | Optional | - |
| `Sessions` | [`[]models.StatsMxtunnelSession`](../../doc/models/stats-mxtunnel-session.md) | Optional | List of sessions<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `State` | [`*models.StatsMxtunnelStateEnum`](../../doc/models/stats-mxtunnel-state-enum.md) | Optional | enum: `established`, `established_with_sessions`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `TxControlPkts` | `*int` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "last_seen": 1470417522.0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "remote_ip": "remote_ip4",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap": "ap4",
  "for_site": false,
  "fwupdate": {
    "progress": 100,
    "status": "upgraded",
    "status_id": 70,
    "timestamp": 147.68,
    "will_retry": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "mtu": 34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

