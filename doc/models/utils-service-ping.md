
# Utils Service Ping

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsServicePing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | **Default**: `10` |
| `Host` | `string` | Required | - |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Service` | `string` | Required | Ping packet takes the same path as the service |
| `Size` | `*int` | Optional | **Default**: `56`<br>**Constraints**: `>= 56`, `<= 65535` |
| `Tenant` | `*string` | Optional | Tenant context in which the packet is sent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "count": 10,
  "host": "host0",
  "service": "service8",
  "size": 56,
  "node": "node0",
  "tenant": "tenant4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

