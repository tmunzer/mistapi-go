
# Stats Mxtunnel Session

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxtunnelSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LocalSid` | `int` | Required | Remote sessions id (dynamically unless Tunnel is said to be static) |
| `RemoteId` | `string` | Required | WxlanTunnel Remote ID |
| `RemoteSid` | `int` | Required | Remote sessions id (dynamically unless Tunnel is said to be static) |
| `State` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "local_sid": 114,
  "remote_id": "remote_id0",
  "remote_sid": 78,
  "state": "state8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

