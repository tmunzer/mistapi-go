
# Utils Show Dhcp Leases

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowDhcpLeases`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Network` | `string` | Required | DHCP network for the leases, returns full table if not specified |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "network": "guest",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

