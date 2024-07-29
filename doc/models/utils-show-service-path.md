
# Utils Show Service Path

The exact service name for which to display the service path

## Structure

`UtilsShowServicePath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "service_name": "any",
  "node": "node0"
}
```

