
# Utils Release Dhcp

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsReleaseDhcp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `PortId` | `string` | Required | The network interface on which to release the current DHCP release<br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/1.10",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

