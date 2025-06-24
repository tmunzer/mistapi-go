
# Utils Ping

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | **Default**: `10` |
| `EgressInterface` | `*string` | Optional | Interface through which packet needs to egress |
| `Host` | `string` | Required | - |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Size` | `*int` | Optional | **Default**: `56`<br><br>**Constraints**: `>= 56`, `<= 65535` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "count": 10,
  "host": "1.1.1.1",
  "size": 56,
  "egress_interface": "egress_interface2",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

