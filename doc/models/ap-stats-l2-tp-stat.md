
# Ap Stats L2 Tp Stat

## Structure

`ApStatsL2TpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Sessions` | [`[]models.ApStatsL2TpStatSession`](../../doc/models/ap-stats-l2-tp-stat-session.md) | Optional | list of sessions |
| `State` | [`*models.L2TpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | - |
| `Uptime` | `models.Optional[int]` | Optional | uptime |
| `WxtunnelId` | `models.Optional[uuid.UUID]` | Optional | WxlanTunnel ID |

## Example (as JSON)

```json
{
  "state": "established",
  "uptime": 135,
  "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
  "sessions": [
    {
      "local_sid": 84,
      "remote_id": "remote_id6",
      "remote_sid": 208,
      "state": "wait-ctrl-reply"
    },
    {
      "local_sid": 84,
      "remote_id": "remote_id6",
      "remote_sid": 208,
      "state": "wait-ctrl-reply"
    },
    {
      "local_sid": 84,
      "remote_id": "remote_id6",
      "remote_sid": 208,
      "state": "wait-ctrl-reply"
    }
  ]
}
```

