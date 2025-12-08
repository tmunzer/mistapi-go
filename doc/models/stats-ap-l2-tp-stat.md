
# Stats Ap L2 Tp Stat

## Structure

`StatsApL2tpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Sessions` | [`[]models.StatsApL2tpStatSession`](../../doc/models/stats-ap-l2-tp-stat-session.md) | Optional | List of sessions |
| `State` | [`*models.L2tpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `Uptime` | `models.Optional[int]` | Optional | Uptime |
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
      "state": "established_with_session"
    }
  ]
}
```

