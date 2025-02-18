
# Utils Show Session

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | The exact service name for which to display the active sessions |
| `SessionId` | `*string` | Optional | Show session details by session_id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "service_name": "any",
  "node": "node0",
  "session_id": "session_id4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

