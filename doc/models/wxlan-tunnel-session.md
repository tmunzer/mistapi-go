
# Wxlan Tunnel Session

## Structure

`WxlanTunnelSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAsSessionId` | `*string` | Optional | if `use_ap_as_session_ids`==`true`, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id |
| `Comment` | `*string` | Optional | optional, user-specified string for display purpose |
| `EnableCookie` | `*bool` | Optional | - |
| `Ethertype` | [`*models.WxlanTunnelSessionEthertypeEnum`](../../doc/models/wxlan-tunnel-session-ethertype-enum.md) | Optional | - |
| `LocalSessionId` | `*int` | Optional | 1-4294967295<br>**Constraints**: `>= 1`, `<= 4294967295` |
| `Pseudo8021adEnabled` | `*bool` | Optional | optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers.<br>**Default**: `false` |
| `RemoteId` | `*string` | Optional | remote-id of the session, has to be unique in the same tunnel |
| `RemoteSessionId` | `*int` | Optional | 1-4294967295<br>**Constraints**: `>= 1`, `<= 4294967295` |
| `UseApAsSessionIds` | `*bool` | Optional | whether to use AP (last 4 bytes of MAC currently) as session ids<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "pseudo_802.1ad_enabled": false,
  "use_ap_as_session_ids": false,
  "ap_as_session_id": "ap_as_session_id2",
  "comment": "comment6",
  "enable_cookie": false,
  "ethertype": "ethernet",
  "local_session_id": 10
}
```
