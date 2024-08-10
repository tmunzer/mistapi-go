
# Ap Stats L2 Tp Stat Session

## Structure

`ApStatsL2tpStatSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `models.Optional[int]` | Optional | remote sessions id (dynamically unless Tunnel is said to be static) |
| `RemoteId` | `models.Optional[string]` | Optional | WxlanTunnel Remote ID (user-configured) |
| `RemoteSid` | `models.Optional[int]` | Optional | remote sessions id (dynamically unless Tunnel is said to be static) |
| `State` | [`*models.L2tpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |

## Example (as JSON)

```json
{
  "local_sid": 31,
  "remote_id": "vpn1",
  "remote_sid": 13,
  "state": "established"
}
```

