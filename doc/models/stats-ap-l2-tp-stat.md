
# Stats Ap L2 Tp Stat

L2TP tunnel status reported by an AP

## Structure

`StatsApL2tpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Sessions` | [`[]models.StatsApL2tpStatSession`](../../doc/models/stats-ap-l2-tp-stat-session.md) | Optional | L2TP tunnel sessions reported by an AP |
| `State` | [`*models.L2tpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `Uptime` | `models.Optional[int]` | Optional, Read-only | Tunnel uptime, in seconds |
| `WxtunnelId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Identifier of the associated WxLAN tunnel |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsApL2tpStat := models.StatsApL2tpStat{
        Sessions:             []models.StatsApL2tpStatSession{
            models.StatsApL2tpStatSession{
                LocalSid:             models.NewOptional(models.ToPointer(84)),
                RemoteId:             models.NewOptional(models.ToPointer("remote_id6")),
                RemoteSid:            models.NewOptional(models.ToPointer(208)),
                State:                models.ToPointer(models.L2tpStateEnum_ESTABLISHEDWITHSESSION),
            },
            models.StatsApL2tpStatSession{
                LocalSid:             models.NewOptional(models.ToPointer(84)),
                RemoteId:             models.NewOptional(models.ToPointer("remote_id6")),
                RemoteSid:            models.NewOptional(models.ToPointer(208)),
                State:                models.ToPointer(models.L2tpStateEnum_ESTABLISHEDWITHSESSION),
            },
            models.StatsApL2tpStatSession{
                LocalSid:             models.NewOptional(models.ToPointer(84)),
                RemoteId:             models.NewOptional(models.ToPointer("remote_id6")),
                RemoteSid:            models.NewOptional(models.ToPointer(208)),
                State:                models.ToPointer(models.L2tpStateEnum_ESTABLISHEDWITHSESSION),
            },
        },
        State:                models.ToPointer(models.L2tpStateEnum_ESTABLISHED),
        Uptime:               models.NewOptional(models.ToPointer(135)),
        WxtunnelId:           models.NewOptional(models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216"))),
    }

}
```

