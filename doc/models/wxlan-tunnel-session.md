
# Wxlan Tunnel Session

L2TPv3 session established inside a WxLAN tunnel

## Structure

`WxlanTunnelSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAsSessionId` | `*string` | Optional | If `use_ap_as_session_ids`==`true`, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id |
| `Comment` | `*string` | Optional | Optional, user-specified string for display purpose |
| `EnableCookie` | `*bool` | Optional | Whether L2TPv3 cookie support is enabled for this session |
| `Ethertype` | [`*models.WxlanTunnelSessionEthertypeEnum`](../../doc/models/wxlan-tunnel-session-ethertype-enum.md) | Optional | Frame type carried by this tunnel session. enum: `ethernet`, `vlan` |
| `LocalSessionId` | `*int` | Optional | Identifier for the local L2TPv3 session, from 1 to 2147483647<br><br>**Constraints**: `>= 1`, `<= 2147483647` |
| `Pseudo8021adEnabled` | `*bool` | Optional | Optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers.<br><br>**Default**: `false` |
| `RemoteId` | `*string` | Optional | Remote-id of the session, has to be unique in the same tunnel |
| `RemoteSessionId` | `*int` | Optional | Identifier for the remote L2TPv3 session, from 1 to 2147483647<br><br>**Constraints**: `>= 1`, `<= 2147483647` |
| `UseApAsSessionIds` | `*bool` | Optional | Whether to use AP (last 4 bytes of MAC currently) as session ids<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wxlanTunnelSession := models.WxlanTunnelSession{
        ApAsSessionId:        models.ToPointer("ap_as_session_id0"),
        Comment:              models.ToPointer("comment2"),
        EnableCookie:         models.ToPointer(false),
        Ethertype:            models.ToPointer(models.WxlanTunnelSessionEthertypeEnum_ETHERNET),
        LocalSessionId:       models.ToPointer(64),
        Pseudo8021adEnabled:  models.ToPointer(false),
        UseApAsSessionIds:    models.ToPointer(false),
    }

}
```

