
# Stats Mxtunnel Session

Session information for a WxLAN or Mist tunnel

## Structure

`StatsMxtunnelSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `int` | Required | Local session identifier for the tunnel session |
| `RemoteId` | `string` | Required | Configured remote identifier for the tunnel session |
| `RemoteSid` | `int` | Required | Peer session identifier for the tunnel session |
| `State` | `string` | Required | Current state of the tunnel session |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxtunnelSession := models.StatsMxtunnelSession{
        LocalSid:             160,
        RemoteId:             "remote_id0",
        RemoteSid:            132,
        State:                "state2",
    }

}
```

