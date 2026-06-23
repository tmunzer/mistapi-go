
# Stats Ap L2 Tp Stat Session

L2TP tunnel session reported by an AP

## Structure

`StatsApL2tpStatSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `models.Optional[int]` | Optional, Read-only | Local L2TP session identifier for this tunnel session |
| `RemoteId` | `models.Optional[string]` | Optional, Read-only | User-configured remote identifier for the WxLAN tunnel |
| `RemoteSid` | `models.Optional[int]` | Optional, Read-only | Remote L2TP session identifier for this tunnel session |
| `State` | [`*models.L2tpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApL2tpStatSession := models.StatsApL2tpStatSession{
        LocalSid:             models.NewOptional(models.ToPointer(31)),
        RemoteId:             models.NewOptional(models.ToPointer("vpn1")),
        RemoteSid:            models.NewOptional(models.ToPointer(13)),
        State:                models.ToPointer(models.L2tpStateEnum_ESTABLISHED),
    }

}
```

