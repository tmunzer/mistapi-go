
# Stats Ap L2 Tp Stat Session

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApL2tpStatSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `models.Optional[int]` | Optional | Remote sessions id (dynamically unless Tunnel is said to be static) |
| `RemoteId` | `models.Optional[string]` | Optional | WxlanTunnel Remote ID (user-configured) |
| `RemoteSid` | `models.Optional[int]` | Optional | Remote sessions id (dynamically unless Tunnel is said to be static) |
| `State` | [`*models.L2tpStateEnum`](../../doc/models/l2-tp-state-enum.md) | Optional | enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "local_sid": 31,
  "remote_id": "vpn1",
  "remote_sid": 13,
  "state": "established",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

