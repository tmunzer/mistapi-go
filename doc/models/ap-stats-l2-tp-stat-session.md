
# Ap Stats L2 Tp Stat Session

## Structure

`ApStatsL2TpStatSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `*int` | Optional | remote sessions id (dynamically unless Tunnel is said to be static) |
| `RemoteId` | `*string` | Optional | WxlanTunnel Remote ID (user-configured) |
| `RemoteSid` | `*int` | Optional | remote sessions id (dynamically unless Tunnel is said to be static) |
| `State` | [`*models.L2TpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "local_sid": 31,
  "remote_id": "vpn1",
  "remote_sid": 13,
  "state": "established"
}
```

