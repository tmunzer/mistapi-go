
# Utils Clear Session

To use five tuples to lookup the session to be cleared, all must be provided

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | Service name, only supported in SSR |
| `SessionIds` | `[]uuid.UUID` | Optional | List of id of the sessions to be cleared |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "service_name": "internet-wan_and_lte",
  "session_ids": [
    "88776655-0123-4567-890a-112233445566"
  ],
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

