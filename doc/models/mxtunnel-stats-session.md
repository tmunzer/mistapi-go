
# Mxtunnel Stats Session

## Structure

`MxtunnelStatsSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `int` | Required | remote sessions id (dynamically unless Tunnel is said to be static) |
| `RemoteId` | `string` | Required | WxlanTunnel Remote ID |
| `RemoteSid` | `int` | Required | remote sessions id (dynamically unless Tunnel is said to be static) |
| `State` | `string` | Required | - |

## Example (as JSON)

```json
{
  "local_sid": 248,
  "remote_id": "remote_id4",
  "remote_sid": 44,
  "state": "state2"
}
```

